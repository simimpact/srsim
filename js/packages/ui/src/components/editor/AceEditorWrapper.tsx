import AceEditor from "react-ace";

//THESE IMPORTS NEEDS TO BE AFTER IMPORTING AceEditor
import "ace-builds/src-min-noconflict/ext-language_tools";
//manually import supported themes cause we can't get for loop to work here
import "ace-builds/src-min-noconflict/theme-github";
import "ace-builds/src-min-noconflict/theme-kuroir";
import "ace-builds/src-min-noconflict/theme-monokai";
import "ace-builds/src-min-noconflict/theme-solarized_dark";
import "ace-builds/src-min-noconflict/theme-solarized_light";
import "ace-builds/src-min-noconflict/theme-terminal";
import "ace-builds/src-min-noconflict/theme-textmate";
import "ace-builds/src-min-noconflict/theme-tomorrow";
import "ace-builds/src-min-noconflict/theme-tomorrow_night";
import "ace-builds/src-min-noconflict/theme-twilight";
import "ace-builds/src-min-noconflict/theme-xcode";
import "ace-builds/src-min-noconflict/mode-yaml";
import { AceEditorWrapperProps } from "./types.js";

export function AceEditorWrapper({
  cfg,
  onChange,
  maxLines = 35,
  fontSize = 14,
  theme = "tomorrow_night",
}: AceEditorWrapperProps) {
  return (
    <AceEditor
      mode="yaml"
      theme={theme}
      width="100%"
      onChange={onChange}
      value={cfg}
      name="config_editor"
      editorProps={{
        $blockScrolling: true,
      }}
      setOptions={{
        maxLines: maxLines,
        fontSize: fontSize,
        tabSize: 2,
        highlightActiveLine: false,
      }}
    />
  );
}
