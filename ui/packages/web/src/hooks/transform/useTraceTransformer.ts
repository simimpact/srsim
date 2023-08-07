export function useTraceTransformer() {
  function toFullTraces(charId: number, shortTraces: string[] | number[]) {
    return shortTraces.map(shortHand => Number(`${charId}${shortHand}`));
  }

  return { toFullTraces };
}
