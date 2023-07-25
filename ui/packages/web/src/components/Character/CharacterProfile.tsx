import { cva } from "class-variance-authority";
import { useRef, useState } from "react";
import { AvatarSkillConfig, SkillType } from "@/bindings/AvatarSkillConfig";
import {
  useCharacterEidolon,
  useCharacterSearch,
  useCharacterSkill,
  useCharacterTrace,
} from "@/hooks/queries/useCharacter";
import { useLightConeSearch } from "@/hooks/queries/useLightCone";
import { useTraceTransformer } from "@/hooks/transform/useTraceTransformer";
import { CharacterConfig } from "@/providers/temporarySimControlTypes";
import { cn } from "@/utils/classname";
import { ImpositionIcon } from "../LightCone/ImpositionIcon";
import { LightConePortrait } from "../LightCone/LightConePortrait";
import { Badge } from "../Primitives/Badge";
import { Button } from "../Primitives/Button";
import { RelicItem } from "../Relic/RelicItem";
import { CharacterFloatCard } from "./CharacterFloatCard";
import { CharacterStatTable } from "./CharacterStatTable";
import { EidolonIcon } from "./EidolonIcon";
import { SkillIcon } from "./SkillIcon";
import { TraceTree } from "./TraceTree";

interface Props {
  data: CharacterConfig;
}
const CharacterProfile = ({ data: configData }: Props) => {
  const { character } = useCharacterSearch(configData.key);
  const { skills } = useCharacterSkill(character?.avatar_id);
  const { eidolons } = useCharacterEidolon(character?.avatar_id);
  const { traces } = useCharacterTrace(character?.avatar_id);
  const { lightCone } = useLightConeSearch(configData.light_cone.key);
  const { toFullTraces } = useTraceTransformer();
  const [open, setOpen] = useState(false);
  const relicHeight = useRef<HTMLDivElement>(null);

  if (!character || !lightCone) return null;

  const configTraces = toFullTraces(character.avatar_id, configData.traces);

  const { light_cone, abilities, eidols } = configData;

  const params: SkillType[] = ["Maze", "Normal", "Talent", "BPSkill", "Ultra"];
  const [technique, basic, talent, skill, ult] = params.map(e => getSkill(skills, e));

  const skillRowVariant = cva("flex flex-col gap-2");

  const skillColumns = [
    { lv: 1, maxLv: 1, cfg: technique, trace: 0 },
    { lv: abilities.attack, maxLv: maxSkillLevel(eidols).attack, cfg: basic, trace: 2 },
    { lv: abilities.talent, maxLv: maxSkillLevel(eidols).talent, cfg: talent, trace: 4 },
    { lv: abilities.skill, maxLv: maxSkillLevel(eidols).skill, cfg: skill, trace: 6 },
    { lv: abilities.ult, maxLv: maxSkillLevel(eidols).ult, cfg: ult, trace: null },
  ];

  return (
    <div id="main-container" className="grid grid-cols-12 gap-2.5">
      <div id="left-container" className="col-span-6 grid grid-cols-12">
        <div id="char-img" className="col-span-3 flex flex-col items-center">
          <CharacterFloatCard
            className="p-6"
            imgUrl={`https://raw.githubusercontent.com/Mar-7th/StarRailRes/master/image/character_preview/${character.avatar_id}.png`}
            {...character}
          />

          <div className="font-bold text-xl flex flex-col justify-center">
            <span className="text-center">{character.avatar_name}</span>
            <span className="text-lg text-center">
              Lv. {configData.level} / {configData.max_level}
            </span>
          </div>
        </div>

        <div id="char-info" className="col-span-9">
          <div id="eidolon-skill-spans" className="flex">
            <div id="eidolon" className="flex flex-col gap-2">
              <Badge className="w-fit self-center">E{configData.eidols}</Badge>
              <div className="flex flex-col gap-1">
                {eidolons?.map(eidolon => (
                  <EidolonIcon
                    key={eidolon.rank}
                    data={eidolon}
                    characterId={character.avatar_id}
                    disabled={eidolon.rank > configData.eidols}
                  />
                ))}
              </div>
            </div>

            <div id="skill-trace" className="grid grid-cols-5">
              {skillColumns.map(({ lv, maxLv, cfg, trace }) => (
                <div key={trace} className={skillRowVariant()}>
                  <Badge className="w-fit self-center">
                    {lv} / {maxLv}
                  </Badge>
                  {cfg && <SkillIcon data={cfg} characterId={character.avatar_id} slv={lv} />}
                  {trace !== null && <Badge className="w-fit self-center">A{trace}</Badge>}
                  {traces && trace !== null && (
                    <TraceTree
                      emptyBigTrace={trace == 0}
                      bigTraceAscension={trace}
                      path={character.avatar_base_type}
                      traces={traces}
                      charTraces={configTraces}
                    />
                  )}
                </div>
              ))}
            </div>
          </div>
        </div>

        <div id="lc-img" className="relative col-span-3">
          {/* alternate version using the 'skewed'/rotated ingame image
          <LightConeCard
            rarity={lightConeMetadata.rarity}
            path={lightConeMetadata.avatar_base_type}
            name={lightConeMetadata.equipment_name}
            imgUrl={lightConeIconUrl(lightConeMetadata.equipment_id)}
          /> */}
          <div className="flex flex-col justify-center items-center">
            <div className="font-bold text-xl flex flex-col justify-center">
              <span className="text-center">{lightCone.equipment_name}</span>

              <div className="text-lg flex gap-2 justify-center items-center">
                <span>
                  Lv. {light_cone.level} / {light_cone.max_level}
                </span>

                <ImpositionIcon imposition={light_cone.imposition} />
              </div>
            </div>

            <div className="p-6">
              <LightConePortrait data={lightCone} />
            </div>
          </div>
        </div>

        {/*
        <LightConeInfo
          id="lc-info"
          className="col-span-9 flex flex-col"
          name={lightCone.equipment_name}
          imposition={light_cone.imposition}
          currentLevel={light_cone.level}
          maxLevel={light_cone.max_level}
        />
        */}
      </div>

      <CharacterStatTable
        id="mid-stats"
        className="col-span-3 grid h-fit grid-cols-2 gap-y-2 rounded-md border p-2"
        style={{ height: relicHeight.current?.clientHeight ?? "auto" }}
      />

      <div id="right-relic" className="col-span-3 flex flex-col gap-2">
        <div id="relic-block" className="relative flex flex-col gap-4">
          <div id="4p" className="flex flex-col gap-2" ref={relicHeight}>
            {configData.relics
              ?.slice(0, 4)
              .map((relic, index) => <RelicItem key={index} data={relic} mockIndex={index} />)}
          </div>

          <div id="2p" className="flex flex-col gap-2">
            {configData.relics
              ?.slice(-2)
              .map((relic, index) => (
                <RelicItem key={index} data={relic} mockIndex={index} asSet />
              ))}
          </div>

          <p
            className={cn(
              "bg-background/80 absolute h-full w-full rounded-md p-4 transition-all duration-500 ease-in-out",
              open ? "opacity-100" : "opacity-0"
            )}
          >
            2pc Set Effect <br />
            Increases Lightning DMG by 10%. <br />
            4pc Set Effect <br />
            When the wearer uses their Skill, increases the wearer
            {"'"}s ATK by 20% for 1 turn(s).
          </p>
        </div>

        <Button onClick={() => setOpen(!open)}>Set Bonuses</Button>
      </div>
    </div>
  );
};

function maxSkillLevel(eidolon: number) {
  let attack = 6;
  let skill = 10;
  let ult = 10;
  let talent = 10;

  if (eidolon >= 3) {
    ult += 2;
    talent += 2;
  }
  if (eidolon >= 5) {
    attack += 1;
    skill += 2;
  }
  return {
    attack,
    skill,
    ult,
    talent,
  };
}

function getSkill(list: AvatarSkillConfig[] | undefined, attackType: SkillType) {
  return list?.find(e =>
    e.attack_type ? e.attack_type == attackType : e.skill_type_desc == "Talent"
  );
}

export { CharacterProfile };
