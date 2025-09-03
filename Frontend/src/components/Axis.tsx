import { Arrow, Group, Line, Text } from "react-konva";
import { Fragment } from "react/jsx-runtime";

interface AxisProps {
  direction: "v" | "h";
  x: number;
  y: number;
  length: number;
  interval: number;
}

export function Axis({ direction, x, y, length, interval }: AxisProps) {
  const mainLinePoints =
    direction === "v" ? [0, 0, 0, -length] : [0, 0, length, 0];

  const tickSize = 10;

  const nTicks = Math.ceil(length / interval);
  return (
    <Group x={x} y={y}>
      <Arrow
        points={mainLinePoints}
        fill={"black"}
        strokeWidth={2}
        stroke={"black"}
      />
      {Array.from({ length: nTicks }).map((_, i) => (
        <Fragment key={i}>
          <Line
            points={
              direction === "v"
                ? [-tickSize, -i * interval, tickSize, -i * interval]
                : [i * interval, -tickSize, i * interval, tickSize]
            }
            stroke={"black"}
            strokeWidth={2}
            key={i}
          />
          {direction === "v" ? (
            <Text
              x={-tickSize - 25}
              y={-i * interval - 8}
              width={30}
              align="left"
              text={String(i * interval)}
              fontSize={12}
              fill="black"
              fontStyle="bold"
            />
          ) : (
            <Text
              x={i * interval - 10}
              y={tickSize + 2}
              width={30}
              align="left"
              text={String(i * interval)}
              fontSize={12}
              fill="black"
              fontStyle="bold"
            />
          )}
        </Fragment>
      ))}
    </Group>
  );
}
