import { atom } from "jotai";

export const itemnameAtom = atom("danheng");

export const ascensionAtom = atom(0);

export const maxLevelAtom = atom(get => get(ascensionAtom) * 10 + 20);
