import { useLocation, useNavigate } from "react-router-dom";

/**
 * this is a helper hook when url affects what tab is selected on screen
 * (tabs that are visited frequently and usually shared via links)
 * @returns tab key of the tab, when passing this to the `<Tabs />` component,
 * please use `value` instead of `defaultValue` and use `??` to denote
 * default Value
 */
export function useTabRouteHelper() {
  const { hash } = useLocation();
  const navigate = useNavigate();
  // slice(1) removes leading #
  const { tab } = Object.fromEntries(new URLSearchParams(hash.slice(1))) as unknown as {
    tab: string | undefined;
  };

  function setTab(toTabKey: string) {
    navigate(`#tab=${toTabKey}`);
  }
  return { tab, setTab };
}
