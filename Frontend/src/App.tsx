import { Layer } from "react-konva";
import { CustomStage } from "./components/CustomStage";
import { useStageDimensions } from "./hooks/useStageDimensions";
import { CartesianPlane } from "./components/CartesianPlane";
import { useEffect } from "react";
import { useEventStream } from "./hooks/userEventStream";
import type { Solution } from "./types/Solution";
import { useTSPState } from "./hooks/userTSPState";

export function App() {
  const { height, width } = useStageDimensions();
  const {
    cities,
    optimizedPath,
    addCity,
    removeLastCity,
    clearCities,
    setOptimization,
    setOptimizedAsPath,
  } = useTSPState();

  const { startEventStream, messages, stopEventStream } =
    useEventStream<Solution>();

  const handleClickSolve = () => {
    stopEventStream();
    startEventStream({
      method: "POST",
      body: { cities },
    });
  };

  useEffect(() => {
    if (messages.length > 0) {
      setOptimization(messages[messages.length - 1].idxs);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [messages]);

  return (
    <div className="h-screen w-screen bg-white flex flex-col">
      <div className="navbar bg-base-100 shadow-sm gap-2">
        <h1 className="text-lg font-bold px-4">TSP Vizu</h1>
        <button onClick={clearCities} className="btn bg-red-500">
          Clear cities
        </button>
        <button
          onClick={() => {
            stopEventStream();
            removeLastCity();
          }}
          className="btn bg-blue-500"
        >
          Remove last city
        </button>
        <button
          onClick={() => {
            stopEventStream();
            setOptimizedAsPath();
          }}
          className="btn bg-blue-500"
        >
          Take current solution
        </button>
        <button onClick={handleClickSolve} className="btn bg-green-500">
          Solve
        </button>
      </div>
      <CustomStage height={height * 0.9} width={width}>
        <Layer>
          <CartesianPlane
            height={height * 0.8}
            width={width * 0.9}
            path={optimizedPath}
            onClick={(p) => {
              stopEventStream();
              addCity(p);
            }}
          />
        </Layer>
      </CustomStage>
    </div>
  );
}
