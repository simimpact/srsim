import { useQuery } from "@tanstack/react-query";
import API from "@/utils/constants";
import { useToastError } from "./useToastError";

export function useRelicSearch(relicName: string | undefined) {
  const query = useQuery({
    queryKey: ["relicSet", relicName],
    queryFn: async () => await API.relicSet.get(relicName),
    enabled: !!relicName,
  });

  useToastError(query);

  return { relicSetConfig: query.data };
}
