import { useQuery } from "@tanstack/react-query";
import API from "@/utils/constants";

export function useLightConeSearch(name: string | undefined) {
  const query = useQuery({
    queryKey: ["lightCone", name],
    queryFn: async () => await API.lightConeSearch.get(name),
  });
  return { lightCone: query.data };
}
