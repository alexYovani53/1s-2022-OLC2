
import CodeEditor from "@uiw/react-textarea-code-editor";
import React, { useRef } from 'react';
import  './css/editor.css'
import Editor from "./view/Editor";
import axios from "axios";

export default function App() {

     const [code, setCode] = React.useState(`public class Main
    {
        public static void main(String[] args) {
            system.out.println("Ejemplo1 " + (5+6--9 + 1));
        }
    }
    `);
  const editor1  = useRef();
  const editor2  = useRef();


  const setFocus =()=>{
    editor1.current.focus();
  }

  const enviarDatos = ()=>{

    console.log(editor1.current.getCode())

    axios.post("http://localhost:3000/prueba",{
      text:editor1.current.getCode()
    })
    .then(result=>{

      editor2.current.establecerCodigo(result.data.val);
    })
    .catch(err=>console.log(err))
  }

  

  return (
    <div className="pagina">
{/* 
          <CodeEditor
            value={code}
            language="js"
            placeholder="Please enter JS code."
            onChange={(evn) => setCode(evn.target.value)}
            padding={15}

            style={{
              fontSize: 12,
              backgroundColor: "#f5f5f5",
              fontFamily: 'ui-monospace,SFMono-Regular,SF Mono,Consolas,Liberation Mono,Menlo,monospace',
            }}
          /> */}
          <div className="container" >
            <Editor  ref={editor1} cod={code} onClick={setFocus}></Editor>
            <button className="enviar" onClick={enviarDatos}> Enviar</button>
            <Editor ref={editor2} cod={""}></Editor>
          </div>
    </div>

  );
}