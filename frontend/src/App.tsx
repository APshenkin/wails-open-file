import {useEffect, useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {Greet} from "../wailsjs/go/main/App";
import {EventsEmit, EventsOn} from "../wailsjs/runtime";

function App() {
  const [resultText, setResultText] = useState("Please enter your name below 👇");
  const [name, setName] = useState('');
  const updateName = (e: any) => setName(e.target.value);
  const updateResultText = (result: string) => setResultText(result);

  function greet() {
    Greet(name).then(updateResultText);
  }

  useEffect(() => {
    // subscribe to file opened event
    EventsOn("fileOpened", (str: string) => updateResultText(`file association support! file is opened ${str}`))
    // subscribe to file opened event
    EventsOn("customUrlOpened", (str: string) => updateResultText(`custom url is opened ${str}`))
    // subscribe to launch Args event
    EventsOn("launchArgs", (str: string[]) => updateResultText(`launch args ${str}`))
    // emit frontEndLoaded event to backend
    EventsEmit("frontEndLoaded")
  }, [])

  return (
    <div id="App">
      <img src={logo} id="logo" alt="logo"/>
      <div id="result" className="result">{resultText}</div>
      <div id="input" className="input-box">
        <input id="name" className="input" onChange={updateName} autoComplete="off" name="input" type="text"/>
        <button className="btn" onClick={greet}>Greet</button>
      </div>
    </div>
  )
}

export default App
