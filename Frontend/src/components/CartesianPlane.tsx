import { Group, Line, Rect, Text } from "react-konva";
import { Axis } from "./Axis";
import type { Point } from "../types/Point";
import type { KonvaEventObject, Node, NodeConfig } from "konva/lib/Node";

interface CartesianPlaneProps {
  height: number;
  width: number;
  path?: Point[];
  onClick?: (p: Point) => void;
}

export function CartesianPlane({
  height,
  width,
  path,
  onClick,
}: CartesianPlaneProps) {
  const handleRectClick = (
    e: KonvaEventObject<MouseEvent, Node<NodeConfig>>
  ) => {
    if (!onClick) return;

    const clickX = e.evt.layerX - width * 0.05;
    const clickY = e.evt.layerY - height * 0.05;

    const cartesianPoint: Point = {
      x: clickX,
      y: height - clickY,
    };

    if (
      clickX >= 0 &&
      clickX <= width &&
      cartesianPoint.y >= 0 &&
      cartesianPoint.y <= height
    ) {
      onClick(cartesianPoint);
    }
  };

  return (
    <Group x={width * 0.05} y={height * 0.05}>
      <Rect
        width={width}
        height={height}
        fill={"white"}
        onClick={handleRectClick}
      />
      <Axis direction="h" x={0} y={height} length={width} interval={100} />
      <Axis direction="v" x={0} y={height} length={height} interval={100} />
      {path && path.length > 1 && (
        <Line
          points={path.flatMap((p) => [p.x, height - p.y])}
          stroke="blue"
          strokeWidth={2}
          closed={false}
        />
      )}
      {path &&
        path.map((p, i) => (
          <Group key={i}>
            <Rect
              x={p.x - 4}
              y={height - p.y - 4}
              width={8}
              height={8}
              fill={i === 0 ? "gold" : "red"}
              stroke={i === 0 ? "orange" : "black"}
              strokeWidth={i === 0 ? 2 : 1}
            />
            <Text
              x={p.x}
              y={height - p.y - 22}
              text={i.toString()}
              fontSize={14}
              fill={i === 0 ? "orange" : "black"}
              fontStyle="bold"
              offsetX={3}
            />
          </Group>
        ))}
    </Group>
  );
}
