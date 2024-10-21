"use client";
import React from "react";
import { AceEditorWrapper } from "./AceEditorWrapper";
import { EditorProps, Theme } from "./types";

const LOCALSTORAGE_THEME_KEY = "gcsim-config-editor-theme";
const LOCALSTORAGE_FONT_SIZE_KEY = "gcsim-config-editor-font-size";

export const Editor = ({ cfg, onChange, className = "" }: EditorProps) => {
  const [theme, setTheme] = React.useState<Theme>("tomorrow_night");
  const [fontSize, setFontSize] = React.useState(14);
  React.useEffect(() => {
    if (typeof window !== "undefined") {
      let theme = window.localStorage.getItem(LOCALSTORAGE_THEME_KEY) ?? "tomorrow_night";
      let fontSize = window.localStorage.getItem(LOCALSTORAGE_FONT_SIZE_KEY)
        ? Number(window.localStorage.getItem(LOCALSTORAGE_FONT_SIZE_KEY))
        : 14;
      setTheme(theme);
      setFontSize(fontSize);
    }
  }, []);
  React.useEffect(() => {
    if (typeof window !== "undefined") {
      window.localStorage.setItem(LOCALSTORAGE_THEME_KEY, theme);
      window.localStorage.setItem(LOCALSTORAGE_FONT_SIZE_KEY, fontSize.toString());
    }
  }, [theme, fontSize]);
  return (
    <div className={className}>
      <AceEditorWrapper cfg={cfg} onChange={onChange} theme={theme} fontSize={fontSize} />
    </div>
  );
};
