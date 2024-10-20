"use client";
import React from "react";
import { AceEditorWrapper } from "./AceEditorWrapper";
import { EditorProps, Theme } from "./types";

const LOCALSTORAGE_THEME_KEY = "gcsim-config-editor-theme";
const LOCALSTORAGE_FONT_SIZE_KEY = "gcsim-config-editor-font-size";

export const Editor = ({ cfg, onChange, className = "" }: EditorProps) => {
  const [theme, setTheme] = React.useState<Theme>(() => {
    return localStorage.getItem(LOCALSTORAGE_THEME_KEY) ?? "tomorrow_night";
  });
  const [fontSize, setFontSize] = React.useState(() => {
    return localStorage.getItem(LOCALSTORAGE_FONT_SIZE_KEY)
      ? Number(localStorage.getItem(LOCALSTORAGE_FONT_SIZE_KEY))
      : 14;
  });
  React.useEffect(() => {
    localStorage.setItem(LOCALSTORAGE_THEME_KEY, theme);
    localStorage.setItem(LOCALSTORAGE_FONT_SIZE_KEY, fontSize.toString());
  }, [theme, fontSize]);
  return (
    <div className={className}>
      <AceEditorWrapper cfg={cfg} onChange={onChange} theme={theme} fontSize={fontSize} />
    </div>
  );
};
