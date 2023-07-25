import { useQuery } from "@tanstack/react-query";
import API from "@/utils/constants";

export function useRelicSearch(relicName: string | undefined) {
  const query = useQuery({
    queryKey: ["relicSet", relicName],
    queryFn: async () => await API.relicSet.get(relicName),
    enabled: !!relicName,
  });
  return { relicSetConfig: query.data };
}
