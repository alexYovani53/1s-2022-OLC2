import React, { useImperativeHandle, useRef } from 'react';
import CodeMirror from '@uiw/react-codemirror';
import { javascript } from '@codemirror/lang-javascript';
import { java } from "@codemirror/lang-java";


const Editor = React.forwardRef((props,referencia) => {


     const [code,setCode] = React.useState(props.cod);


    useImperativeHandle(
        referencia,
        () => ({
            getCode() {
                return code;
            },
            establecerCodigo(codigo){
                setCode(codigo);
            }
        }),
    )

    return (
            <CodeMirror
                value={code}
                height="450px"
                theme={'dark'}
                extensions={[java()]}
                onChange={(value, viewUpdate) => {
                    setCode(value);
                    referencia.current.value = code;
                }}
            />
    );
})

export default Editor;
