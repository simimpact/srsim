import { useQuery } from "@tanstack/react-query";
import API from "@/utils/constants";

export function useCharacterSkill(characterId: number | undefined) {
  const query = useQuery({
    queryKey: ["skill", characterId],
    queryFn: async () => await API.skillsByCharId.get(characterId),
    enabled: !!characterId,
  });
  return { skills: query.data?.list };
}

export function useCharacterEidolon(characterId: number | undefined) {
  const query = useQuery({
    queryKey: ["eidolon", characterId],
    queryFn: async () => await API.eidolon.get(characterId),
    enabled: !!characterId,
  });
  return { eidolons: query.data?.list };
}

export function useCharacterTrace(characterId: number | undefined) {
  const query = useQuery({
    queryKey: ["trace", characterId],
    queryFn: async () => await API.trace.get(characterId),
    enabled: !!characterId,
  });
  return { traces: query.data?.list };
}

export function useCharacterSearch(name: string | undefined) {
  const query = useQuery({
    queryKey: ["character", name],
    queryFn: async () => await API.characterSearch.get(name),
    enabled: !!name,
  });

  return { character: query.data };
}
