import { useQuery } from "@tanstack/react-query";
import API from "@/utils/constants";
import { useToastError } from "./useToastError";

export function useCharacterSkill(characterId: number | undefined) {
  const query = useQuery({
    queryKey: ["skill", characterId],
    queryFn: async () => await API.skillsByCharId.get(characterId),
    enabled: !!characterId,
  });

  useToastError(query);

  return { skills: query.data?.list };
}

export function useCharacterEidolon(characterId: number | undefined) {
  const query = useQuery({
    queryKey: ["eidolon", characterId],
    queryFn: async () => await API.eidolon.get(characterId),
    enabled: !!characterId,
  });

  useToastError(query);

  return { eidolons: query.data?.list };
}

export function useCharacterTrace(characterId: number | undefined) {
  const query = useQuery({
    queryKey: ["trace", characterId],
    queryFn: async () => await API.trace.get(characterId),
    enabled: !!characterId,
  });

  useToastError(query);

  return { traces: query.data?.list };
}

export function useCharacterSearch(name: string | undefined) {
  const query = useQuery({
    queryKey: ["character", name],
    queryFn: async () => await API.characterSearch.get(),
    enabled: !!name,
  });

  useToastError(query);

  return { character: query.data };
}
