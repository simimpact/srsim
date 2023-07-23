import { useQuery } from "@tanstack/react-query";
import { CharacterConfig } from "@/providers/temporarySimControlTypes";
import { cn } from "@/utils/classname";
import API, { characterIconUrl } from "@/utils/constants";
import { range } from "@/utils/helpers";
import { elementVariants, rarityVariants } from "@/utils/variants";
import { LightConePortrait } from "../LightCone/LightConePortrait";

interface Props {
  data: CharacterConfig;
}
const CharacterProfile = ({ data: configData }: Props) => {
  const { data: characterMetadata } = useQuery({
    queryKey: ["character", configData.key],
    queryFn: async () => await API.characterSearch.get(configData.key),
  });

  const { data: lightConeMetadata } = useQuery({
    queryKey: ["lightCone", configData.light_cone.key],
    queryFn: async () => await API.lightConeSearch.get(configData.light_cone.key),
  });

  if (!characterMetadata || !lightConeMetadata) return null;

  const element = characterMetadata.damage_type;

  const { light_cone, abilities, eidols } = configData;

  return (
    <div id="main-container" className="flex gap-2.5">
      <div id="left-container" className="grid grid-cols-12">
        <div id="char-img" className="col-span-3 flex flex-col items-center">
          <img
            src={characterIconUrl(characterMetadata.avatar_id)}
            alt={characterMetadata.avatar_name}
            className={cn(
              "h-32 w-32 rounded-full box-content p-1 border-2",
              elementVariants({ border: element }),
              rarityVariants({ rarity: characterMetadata.rarity as 1 | 2 | 3 | 4 | 5 | null })
            )}
          />
          <div>
            {characterMetadata.avatar_name} - {characterMetadata.rarity} ✦
          </div>
          <div>
            {characterMetadata.damage_type} {characterMetadata.avatar_base_type}
          </div>
        </div>

        <div id="char-info" className="col-span-9">
          <span id="level">
            Lv. {configData.level} / {configData.max_level}
          </span>

          <div id="eidolon-skill-spans" className="flex">
            <div id="eidolon" className="flex flex-col">
              {Array.from(range(1, 6, 1)).map(index => (
                <span key={index}>
                  {index} {configData.eidols >= index ? "true" : "false"}
                </span>
              ))}
            </div>

            <div id="skill-trace" className="grid grid-cols-5">
              <div className="flex flex-col">
                <div>1 / 1</div>
                <div>technique</div>
                <div>A0</div>
              </div>
              <div className="flex flex-col">
                <div>
                  {abilities.attack} / {maxSkillLevel(eidols).attack}
                </div>
                <div>basic</div>
                <div>A2</div>
              </div>
              <div className="flex flex-col">
                <div>
                  {abilities.talent} / {maxSkillLevel(eidols).talent}
                </div>
                <div>talent</div>
                <div>A4</div>
              </div>
              <div className="flex flex-col">
                <div>
                  {abilities.skill} / {maxSkillLevel(eidols).skill}
                </div>
                <div>skill</div>
                <div>A6</div>
              </div>
              <div>
                <div>
                  {abilities.ult} / {maxSkillLevel(eidols).ult}
                </div>
                <div>ultra</div>
              </div>
            </div>
          </div>
        </div>

        <div id="lc-img" className="col-span-3 relative">
          {/* alternate version using the 'skewed'/rotated ingame image
          <LightConeCard
            rarity={lightConeMetadata.rarity}
            path={lightConeMetadata.avatar_base_type}
            name={lightConeMetadata.equipment_name}
            imgUrl={lightConeIconUrl(lightConeMetadata.equipment_id)}
          /> */}
          <div className="p-6">
            <LightConePortrait data={lightConeMetadata} />
          </div>
        </div>

        <div id="lc-info" className="col-span-9 flex flex-col">
          <span>{light_cone.key}</span>
          <span>
            {light_cone.imposition} Lv. {light_cone.level} / {light_cone.max_level}
          </span>
        </div>
      </div>

      <div id="mid-stats">mid-stats</div>

      <div id="right-relic" className="flex flex-col">
        <div id="4p" className="flex flex-col">
          <div>piece1</div>
          <div>piece2</div>
          <div>piece3</div>
          <div>piece4</div>
        </div>

        <div id="2p" className="flex flex-col">
          <div>piece1</div>
          <div>piece2</div>
        </div>
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

export { CharacterProfile };
