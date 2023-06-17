import { useTranslation } from "react-i18next";
import { IconContext } from "react-icons";
import { AiFillGithub } from "react-icons/ai";
import { FaDiscord } from "react-icons/fa";
import { SiKofi } from "react-icons/si";
import { cn } from "@/utils/classname";

export const Footer = () => {
  const { t } = useTranslation();

  const divider = cn(
    "before:block",
    "before:h-px",
    "before:bg-gradient-to-r",
    "before:from-transparent",
    "before:via-white/50"
  );

  const linkClass = cn("flex gap-2 items-center", "!text-gray-400 hover:!text-[#8abbff]");

  return (
    <div className={`w-full ${divider}`}>
      <div className="px-5 xs:px-16 py-3 flex justify-center gap-2 2xl:mx-auto 2xl:container">
        <div className="self-center text-right text-gray-400 text-xs grow shrink-0 w-2/3 max-w-fit">
          {t("blah")}
        </div>
        <div className="m-[5px] border-x-[1px] border-[#fff3]" />
        <div className="flex flex-wrap gap-x-4 gap-y-2 text-lg font-medium shrink grow-0">
          <IconContext.Provider value={{ size: "32px", color: "inherit" }}>
            <a
              className={linkClass}
              href="https://discord.gg/m7jvjdxx7q"
              target="_blank"
              rel="noreferrer"
            >
              <FaDiscord />
              <span>Discord</span>
            </a>
            <a
              className={linkClass}
              href="https://github.com/simimpact/srsim"
              target="_blank"
              rel="noreferrer"
            >
              <AiFillGithub />
              <span>GitHub</span>
            </a>
            <a
              className={linkClass}
              href="https://ko-fi.com/srliao"
              target="_blank"
              rel="noreferrer"
            >
              <SiKofi />
              <span>Ko-Fi</span>
            </a>
          </IconContext.Provider>
        </div>
      </div>
    </div>
  );
};
