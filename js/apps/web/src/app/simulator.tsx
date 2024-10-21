'use client';
import React, {useRef} from 'react';
import {Button, Editor} from '@ui/components';
import {Executor, ExecutorSupplier, ServerExecutor} from '@srsim/executor';
import {useRouter} from 'next/navigation';
import {throttle} from 'lodash-es';
import {model} from '@srsim/ts-types';
import {ViewerContext} from './viewer/provider';

let exec: ServerExecutor | undefined;
const urlKey = 'server-mode-url';
const defaultURL = 'http://127.0.0.1:54321';

export default function Simulator() {
  const [url, setURL] = React.useState<string>((): string => {
    const saved = localStorage.getItem(urlKey);
    if (saved === null) {
      localStorage.setItem(urlKey, defaultURL);
      return defaultURL;
    }
    return saved;
  });
  React.useEffect(() => {
    localStorage.setItem(urlKey, url);
  }, [url]);

  React.useEffect(() => {
    if (exec != null) {
      exec.set_url(url);
    }
  }, [url]);
  const supplier = useRef<ExecutorSupplier<ServerExecutor>>(() => {
    if (exec == null) {
      exec = new ServerExecutor(url);
    }
    return exec;
  });

  return <SimulatorCore exec={supplier.current} />;
}

type SimulatorCoreProps = {
  exec: ExecutorSupplier<Executor>;
};

const DEFAULT_VIEWER_THROTTLE = 100;
// TODO: this should be an user option
const DEFAULT_ITERS = 1000;
const cfgKey = 'user_local_cfg';

const SimulatorCore = ({exec}: SimulatorCoreProps) => {
  const router = useRouter();
  const [cfg, setCfg] = React.useState<string>((): string => {
    const saved = localStorage.getItem(cfgKey);
    if (saved === null) {
      localStorage.setItem(cfgKey, '');
      return '';
    }
    return saved;
  });
  React.useEffect(() => {
    localStorage.setItem(cfgKey, cfg);
  }, [cfg]);
  const {dispatch} = React.useContext(ViewerContext);

  const run = () => {
    const updateResult = throttle(
      (res: model.SimResult, iters: number, hash: string) => {
        console.log('updating result', res);
        dispatch({
          type: 'SET_RESULT',
          payload: {
            result: res,
            progress: (100 * iters) / DEFAULT_ITERS,
            done: iters === DEFAULT_ITERS,
          },
        });
      },
      DEFAULT_VIEWER_THROTTLE,
      {leading: true, trailing: true},
    );

    exec()
      .run(cfg, DEFAULT_ITERS, updateResult)
      .catch((err) => {
        dispatch({
          type: 'SET_ERROR',
          payload: {
            error: err,
            config: cfg,
          },
        });
      });

    router.push('/viewer');
  };
  return (
    <div className="m-3">
      <Editor cfg={cfg} onChange={(v) => setCfg(v)} className="mb-2"></Editor>
      <div className="sticky bottom-0 flex flex-col gap-y-1 z-10">
        <Button variant="secondary" onClick={() => run()}>
          Run
        </Button>
      </div>
    </div>
  );
};
