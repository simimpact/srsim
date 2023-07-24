import { useQuery } from "@tanstack/react-query";
import { useContext, useState } from "react";
import { AvatarConfig } from "@/bindings/AvatarConfig";
import { CharacterEidolon } from "@/components/Character/CharacterEidolon";
import { CharacterProfile } from "@/components/Character/CharacterProfile";
import { CharacterSkill } from "@/components/Character/CharacterSkill";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/Primitives/Tabs";
import { SimControlContext } from "@/providers/SimControl";
import { CharacterConfig } from "@/providers/temporarySimControlTypes";
import API from "@/utils/constants";
import { CharacterLineup } from "../Root/CharacterLineup";
import { SimActionBar } from "../Root/SimActionBar";

const Config = () => {
  const { simulationConfig } = useContext(SimControlContext);
  const [currentCharacter, setCurrentCharacter] = useState<AvatarConfig | undefined>(undefined);
  const [currentCharacterConfig, setCurrentCharacterConfig] = useState<CharacterConfig | undefined>(
    simulationConfig?.characters[0] ? simulationConfig.characters[0] : undefined
  );
  const characterId = currentCharacter?.avatar_id;

  const { data: characterEidolons } = useQuery({
    queryKey: ["eidolon", characterId],
    queryFn: async () => await API.eidolon.get(characterId),
    enabled: !!characterId,
  });

  const { data: characterSkills } = useQuery({
    queryKey: ["skillsByCharId", characterId],
    queryFn: async () => await API.skillsByCharId.get(characterId),
    enabled: !!characterId,
  });

  function onCharacterSelect(charData: AvatarConfig, index: number) {
    setCurrentCharacter(charData);
    if (simulationConfig?.characters[index]) {
      setCurrentCharacterConfig(simulationConfig.characters[index]);
    }
  }

  return (
    <div id="dev" className="flex h-full self-start grow">
      <div className="flex flex-col grow gap-2">
        <div className="flex gap-4 justify-center mx-8">
          <CharacterLineup onCharacterSelect={onCharacterSelect} />
        </div>

        <div className="flex gap-4 px-4">
          <SimActionBar />
          <div className="grow text-accent-foreground flex flex-col rounded-md">
            <Tabs defaultValue="profile">
              <TabsList className="w-full">
                <TabsTrigger value="profile" className="w-full">
                  Profile
                </TabsTrigger>
                <TabsTrigger value="skills" className="w-full">
                  Skills/Eidolons
                </TabsTrigger>
                <TabsTrigger value="relic" className="w-full">
                  Light Cone/Relic
                </TabsTrigger>
                <TabsTrigger value="trace" className="w-full">
                  Traces
                </TabsTrigger>
              </TabsList>

              <div className="bg-muted rounded-md px-4 p-2 my-2">
                <TabsContent value="profile">
                  {currentCharacterConfig && <CharacterProfile data={currentCharacterConfig} />}
                </TabsContent>
                <TabsContent value="relic">todo</TabsContent>
                <TabsContent value="skills">
                  {currentCharacter && characterSkills?.list && characterEidolons?.list && (
                    <div className="flex flex-col">
                      <p>Skill</p>
                      <CharacterSkill
                        skills={characterSkills.list}
                        characterId={currentCharacter.avatar_id}
                        maxEnergy={currentCharacter.spneed}
                      />
                      <p>Eidolons</p>
                      <CharacterEidolon
                        data={characterEidolons.list}
                        characterId={currentCharacter.avatar_id}
                      />
                    </div>
                  )}
                </TabsContent>
                <TabsContent value="trace">todo</TabsContent>
              </div>
            </Tabs>
          </div>
        </div>
      </div>
    </div>
  );
};
export { Config };
