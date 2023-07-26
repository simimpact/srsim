import { UseQueryResult } from "@tanstack/react-query";
import { useEffect } from "react";
import { useToast } from "@/components/Primitives/Toast/useToast";

/**
 * this is a helper hook to fire toast when a tanstack query returns an error
 * @param query tanstack query hook
 */
export function useToastError<T>(query: UseQueryResult<T>) {
  const { toast } = useToast();

  useEffect(() => {
    if (query.error) {
      toast({ title: "Error", description: query.error as string });
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [query.error]);
}
