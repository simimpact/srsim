import { MouseEventHandler, useRef } from "react";

export default function useCardEffect() {
  let bounds: DOMRect | undefined = undefined;
  const flowRef = useRef<HTMLDivElement>(null);
  const glowRef = useRef<HTMLDivElement>(null);

  const rotateToMouse: MouseEventHandler<HTMLDivElement> = e => {
    bounds = flowRef.current?.getBoundingClientRect();
    const mouseX = e.clientX;
    const mouseY = e.clientY;
    const leftX = mouseX - (bounds?.x ?? 0);
    const topY = mouseY - (bounds?.y ?? 0);
    const center = {
      x: leftX - (bounds?.width ?? 0) / 2,
      y: topY - (bounds?.height ?? 0) / 2,
    };
    const distance = Math.sqrt(center.x ** 2 + center.y ** 2);

    if (flowRef.current) {
      flowRef.current.style.transform = `
      scale3d(1.07, 1.07, 1.07)
      rotate3d(
        ${center.y / 100},
        ${-center.x / 100},
        0,
        ${Math.log(distance) * 2}deg
      )
    `;
    }

    if (glowRef.current) {
      glowRef.current.style.backgroundImage = `
      radial-gradient(
        circle at
        ${center.x * 2 + (bounds?.width ?? 0) / 2}px
        ${center.y * 2 + (bounds?.height ?? 0) / 2}px,
        #ffffff55,
        #0000000f
      )
    `;
    }
  };
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const removeListener: MouseEventHandler<HTMLDivElement> = _ => {
    if (flowRef.current) {
      flowRef.current.style.transform = "";
      flowRef.current.style.background = "";
    }
  };
  return { flowRef, glowRef, rotateToMouse, removeListener };
}
