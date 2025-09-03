import { useState, useEffect, useRef, useCallback } from "react";

interface EventStreamOptions {
  method?: "GET" | "POST";
  body?: unknown;
  headers?: Record<string, string>;
}

export function useEventStream<T>() {
  const [messages, setMessages] = useState<T[]>([]);
  const [isConnected, setIsConnected] = useState(false);
  const abortControllerRef = useRef<AbortController | null>(null);
  const readerRef = useRef<ReadableStreamDefaultReader | null>(null);

  const cleanup = useCallback(() => {
    if (abortControllerRef.current) {
      abortControllerRef.current.abort();
      abortControllerRef.current = null;
    }
    if (readerRef.current) {
      readerRef.current.cancel();
      readerRef.current = null;
    }
    setIsConnected(false);
  }, []);

  const startEventStream = useCallback(
    async (options: EventStreamOptions = {}) => {
      cleanup();
      setMessages([]);

      try {
        const controller = new AbortController();
        abortControllerRef.current = controller;

        const { method = "GET", body, headers = {} } = options;

        const requestOptions: RequestInit = {
          method,
          headers: {
            Accept: "text/event-stream",
            "Cache-Control": "no-cache",
            ...headers,
          },
          signal: controller.signal,
        };

        if (method === "POST" && body) {
          requestOptions.body = JSON.stringify(body);
          requestOptions.headers = {
            ...requestOptions.headers,
            "Content-Type": "application/json",
          };
        }

        const response = await fetch(
          import.meta.env.VITE_API_URL + "/solve",
          requestOptions
        );

        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }

        if (!response.body) {
          throw new Error("Response body is null");
        }

        setIsConnected(true);

        const reader = response.body.getReader();
        readerRef.current = reader;
        const decoder = new TextDecoder();

        let buffer = "";
        while (true) {
          const { done, value } = await reader.read();
          if (done) break;

          buffer += decoder.decode(value, { stream: true });
          const lines = buffer.split("\n");

          buffer = lines.pop() || "";
          for (const line of lines) {
            if (line.trim() === "") continue;

            if (line.startsWith("data:")) {
              try {
                const data = line.slice(5);
                if (data === "[DONE]") {
                  cleanup();
                  return;
                }

                const parsed = JSON.parse(data);
                setMessages((prev) => [...prev, parsed]);
              } catch (parseError) {
                console.error("Error parsing SSE data:", parseError);
              }
            } else if (line.startsWith("event: ")) {
              const eventType = line.slice(7);
              console.log("Event type:", eventType);
            }
          }
        }
      } catch (err) {
        if (err instanceof Error && err.name === "AbortError") {
          console.log("Stream aborted");
        } else {
          console.error("Failed to create stream:", err);
        }
        setIsConnected(false);
      }
    },
    [cleanup]
  );

  const stopEventStream = useCallback(() => {
    cleanup();
  }, [cleanup]);

  useEffect(() => {
    return cleanup;
  }, [cleanup]);

  return {
    messages,
    isConnected,
    startEventStream,
    stopEventStream,
  };
}
