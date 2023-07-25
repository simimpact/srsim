import { useQuery } from "@tanstack/react-query";
import API from "@/utils/constants";
import { useToastError } from "./useToastError";

export function useLightConeSearch(name: string | undefined) {
  const query = useQuery({
    queryKey: ["lightCone", name],
    queryFn: async () => await API.lightConeSearch.get(name),
  });

  useToastError(query);

  return { lightCone: query.data };
}
