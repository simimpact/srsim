import { AnimatePresence, motion } from "framer-motion";
import { useState } from "react";
import { CharacterLineup } from "./CharacterLineup";
import { SimActionBar } from "./SimActionBar";

const { div: Mdiv } = motion;
const Root = () => {
  const [enemyOpen, setEnemyOpen] = useState(false);

  return (
    <div className="flex h-full self-start grow">
      <SimActionBar />
      <div className="flex flex-col grow gap-4">
        <div className="flex gap-4 justify-center mx-8">
          <CharacterLineup />
          {/* NOTE: spaghetti framer code*/}
          <Mdiv layoutId="enemyHead" onClick={() => setEnemyOpen(true)} className="cursor-pointer">
            <CharacterLineup isEnemy />
          </Mdiv>
          <AnimatePresence>
            {enemyOpen && (
              <Mdiv
                layoutId="enemyHead"
                className="flex flex-col gap-2 fixed cursor-pointer"
                onClick={() => setEnemyOpen(false)}
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                exit={{ opacity: 0 }}
              >
                <CharacterLineup isEnemy header="wave1" />
                <CharacterLineup isEnemy header="wave2" />
                <CharacterLineup isEnemy header="wave3" />
              </Mdiv>
            )}
          </AnimatePresence>
        </div>
        <div className="bg-slate-500 flex h-full mx-8 rounded-md">
          <p>
            above list: <br />
            left is players team, right is enemy team. <br />
            clicking on enemy side brings up wave list <br />
            {"<<<-"} these buttons should still be consolidated/grouped
          </p>
        </div>
      </div>
    </div>
  );
};
export { Root };
