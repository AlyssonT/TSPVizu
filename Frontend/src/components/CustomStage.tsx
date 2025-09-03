import { Stage, type StageProps } from "react-konva";

export function CustomStage({ ...props }: StageProps) {
  return <Stage {...props} />;
}
