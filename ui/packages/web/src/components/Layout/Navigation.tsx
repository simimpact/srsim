import { useTranslation } from "react-i18next";
import { Link } from "react-router-dom";
import { cn } from "@/utils/classname";

export const Navigation = () => {
  // @footer
  const { t } = useTranslation();

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
      <div className="px-5 xs:px-16 py-3 flex justify-center gap-2 2xl:mx-auto 2xl:container">
        <div className="self-center text-right text-gray-400 text-xs grow shrink-0 w-2/3 max-w-fit">
          {t("blah")}
        </div>
        <div className="m-[5px] border-x-[1px] border-[#fff3]" />
        <div className="flex flex-wrap gap-x-4 gap-y-2 text-lg font-medium shrink grow-0">
          <img src="/images/favicon.ico" alt="SRSim" />
          <Link to="/hello">hello</Link>
          <Link to="/world">world</Link>
        </div>
      </div>
    </div>
  );
};
