import axios from "axios";
import * as qs from "qs";
import { React, useState } from "react";

function App() {
  const [text, setText] = useState("input your text here");
  const [desc, setDesc] = useState("and description here");
  const [response, setResponse] = useState("");
  const [data, setData] = useState();

  let endpoint = "http://localhost:5000/texts";

  async function getAllData() {
    const r = await axios
      .get(endpoint, {
        headers: {
          "Content-type": "application/json; charset=UTF-8",
        },
      })
      .then((r) => setData(r.data))
      .catch((r) => console.error(r));
  }

  return (
    <div className="App">
      <input
        defaultValue={text}
        inputMode="text"
        onChange={(e) => {
          e.preventDefault();
          setText(e.target.value);
        }}
      ></input>
      <input
        defaultValue={desc}
        inputMode="text"
        onChange={(e) => {
          e.preventDefault();
          setDesc(e.target.value);
        }}
      ></input>
      <button
        title="send"
        onClick={async (e) => {
          const r = await axios
            .post(
              endpoint,
              {
                value: text,
                description: desc,
              },
              {
                headers: {
                  "Content-type": "application/json; charset=UTF-8",
                },
              }
            )
            .then((r) => setResponse(r.data))
            .catch((r) => console.error(r));
        }}
      >
        send to api
      </button>
      <div>{JSON.stringify(response)}</div>

      <button
        onClick={async (e) => {
          getAllData();
        }}
      >
        retrieve all data
      </button>
      <div>{JSON.stringify(data)}</div>

      <button
        onClick={async (e) => {
          await getAllData();
          let idss = data.map((x) => Number.parseInt(x.Id));
          let r = await axios
            .delete(endpoint, {
              data: { ids: idss },
              headers: {
                "Content-type": "application/json; charset=UTF-8",
              },
            })
            .catch((e) => console.error(e))
            .then(async (d) => await getAllData());
          console.log(r);
        }}
      >
        remove all data
      </button>
    </div>
  );
}

export default App;
