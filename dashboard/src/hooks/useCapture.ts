import { Dispatch, SetStateAction, useCallback, useEffect, useMemo, useReducer, useRef, useState } from 'react';
import { message } from 'antd';

export enum CaptureStatus {
  WAITING = 'waiting',
  LOADING = 'loading',
  SUCCESS = 'success',
  FAILURE = 'failure',
  CANCEL = 'cancel',
}

enum CaptureAction {
  start,
  success,
  failure,
  cancel,
}

interface CaptureOption<R = unknown> {
  catchError?: boolean; // * Handle the exception or not
  errorPrompt?: (err: Error) => void;

  onSuccess?: (data: R) => void;
  onError?: (err: Error) => void;
  callback?: () => void;
}

interface CaptureState<R = unknown, Q = unknown> {
  status: CaptureStatus;
  args: Q;
  data: R;
  loading: boolean;
}

interface UseCapture<S = unknown, Q = unknown> {
  fn: (args: Q) => Promise<S>;
  initialState: S;
  initialArgs?: Q;
  showError?: boolean;
  immediately?: boolean;
  option?: CaptureOption;
}

const defaultOption: Partial<CaptureOption<unknown>> = {
  catchError: true,
};

const createCapture = (globalOption?: CaptureOption) => {
  const global = { ...defaultOption, ...globalOption };

  const reducerHandle = (
    prevState: CaptureState,
    action: { type: CaptureAction; data: Partial<CaptureState> },
  ): CaptureState => {
    switch (action.type) {
      case CaptureAction.start:
        return { ...prevState, status: CaptureStatus.LOADING, loading: true, args: action.data.args };
      case CaptureAction.success:
        return prevState.status === CaptureStatus.LOADING
          ? { ...prevState, status: CaptureStatus.SUCCESS, loading: false, data: action.data.data }
          : prevState;
      case CaptureAction.failure:
        return prevState.status === CaptureStatus.LOADING
          ? { ...prevState, status: CaptureStatus.FAILURE, loading: false }
          : prevState;
      case CaptureAction.cancel:
        return prevState.status === CaptureStatus.LOADING
          ? { ...prevState, status: CaptureStatus.CANCEL, loading: false }
          : prevState;
      default:
        return prevState;
    }
  };

  const useCapture = <R = unknown, Q extends object = object>(
    opt: UseCapture<R, Q>,
  ): [CaptureState<R, Q>, Dispatch<SetStateAction<Q>>] => {
    const { fn, initialState = null, initialArgs, showError = true, immediately = false, option } = opt;
    const firstly = useRef<boolean>(immediately);
    const _option = useMemo(() => (option ? { ...global, ...option } : global), [option]);

    const [args, setArgs] = useState<Q>(initialArgs!);
    const [state, dispatch] = useReducer(reducerHandle, {
      loading: false,
      args: initialArgs,
      data: initialState,
      status: CaptureStatus.WAITING,
    });

    const _fn = useCallback(fn, []);
    useEffect(() => {
      if (args === undefined) return;
      // * Called when the hooks is firstly mounted
      if (!firstly.current) {
        firstly.current = true;
        return;
      }

      dispatch({ type: CaptureAction.start, data: { args: args } });
      const promise: Promise<R> = _fn.call(null, args);
      promise
        .then(data => {
          dispatch({ type: CaptureAction.success, data: { data } });
          if (_option.onSuccess) _option.onSuccess(data);
        })
        .catch(err => {
          if (showError) _option.errorPrompt && _option.errorPrompt(err);
          dispatch({ type: CaptureAction.failure, data: {} });
          if (_option.onError) _option.onError(err);
          if (!_option.catchError) throw err;
        })
        .finally(_option.callback);
      return () => {
        if (promise.abort) {
          promise.abort();
          dispatch({ type: CaptureAction.cancel, data: {} });
        }
      };
    }, [_fn, _option, args, showError]);

    return [state as CaptureState<R, Q>, setArgs];
  };
  return useCapture;
};

export const useCapture = createCapture({
  errorPrompt: err => message.error(err.message),
  onError: err => console.error('err:', err),
});
