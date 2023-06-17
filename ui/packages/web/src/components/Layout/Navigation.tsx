import { Moon, Sun } from "lucide-react";
import { useTranslation } from "react-i18next";
import { Link } from "react-router-dom";
import useLocalStorage from "use-local-storage";
import { cn } from "@/utils/classname";
import { Button } from "../Primitives/Button";

export const Navigation = () => {
  // @footer
  const { t } = useTranslation();
  const defaultDark = window.matchMedia("(prefers-color-scheme: dark)").matches;
  const [theme, setTheme] = useLocalStorage("theme", defaultDark ? "dark" : "light");
  function toggleTheme() {
    setTheme(theme === "light" ? "dark" : "light");
    const html = document.getElementsByTagName("html")[0];
    html.setAttribute("class", theme);
  }

  // TODO: style later, too lazy
  return (
    <div
      className={cn(
        "w-full",
        "after:block",
        "after:h-px",
        "after:bg-gradient-to-r",
        "after:from-transparent",
        "after:via-white/50"
      )}
    >
      <div className="px-5 xs:px-16 py-3 flex gap-2 2xl:mx-auto">
        <img src="/images/favicon.ico" alt="SRSim" />
        <div className="m-[5px] border-x-[1px] border-[#fff3]" />
        <div className="flex gap-x-4 gap-y-2 text-lg text-foreground font-medium shrink grow-0 items-center w-full">
          <Link to="/hello">Simulator</Link>
          <Link to="/world">Teams DB</Link>
          <Link to="/world">Docs</Link>
          <div className="grow" />
          <Button variant="ghost" size="sm" onClick={toggleTheme}>
            <Sun className="rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0" />
            <Moon className="absolute rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100" />
            <span className="sr-only">Toggle theme</span>
          </Button>
        </div>
      </div>
    </div>
  );
};
