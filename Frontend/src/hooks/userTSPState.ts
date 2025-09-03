import { useState, useMemo } from "react";
import type { Point } from "../types/Point";

export function useTSPState() {
  const [cities, setCities] = useState<Point[]>([]);
  const [optimizedIndices, setOptimizedIndices] = useState<number[]>([]);

  const optimizedPath = useMemo(() => {
    if (optimizedIndices.length === 0) return cities;
    return optimizedIndices.map((idx) => cities[idx]);
  }, [cities, optimizedIndices]);

  const addCity = (city: Point) => {
    setCities((prev) => [...prev, city]);
    setOptimizedIndices([]);
  };

  const removeLastCity = () => {
    setCities((prev) => prev.slice(0, -1));
    setOptimizedIndices([]);
  };

  const clearCities = () => {
    setCities([]);
    setOptimizedIndices([]);
  };

  const setOptimization = (indices: number[]) => {
    setOptimizedIndices(indices);
  };

  const setOptimizedAsPath = () => {
    if (optimizedIndices.length === 0) return;
    setCities((prev) => optimizedIndices.map((idx) => prev[idx]));
    setOptimizedIndices([]);
  };

  return {
    cities,
    optimizedPath,
    addCity,
    removeLastCity,
    clearCities,
    setOptimization,
    setOptimizedAsPath,
  };
}
