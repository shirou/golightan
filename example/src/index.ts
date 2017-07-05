import * as m from "mithril";

import { ClassComponent, CVnode } from 'mithril'

const targets = [
  "(target)",
  "xml",
  "graphql",
  "css",
  "c",
  "json",
];

var state = {
    index: 0,
    target: "",
    src: "",
    result: "",
    setIndex: (index: number) => {
    state.index = index
    state.target = targets[index]
  }
}

const render = (t: string, src: string) => {
    if (t === "" || t === targets[0] || !src || src === "") {
        return;
    }
    m.request({
        url: "http://localhost:8080/render?target=" + t,
        method: "POST",
        data: src,
        deserialize: (value: string) => {return value},
    })
        .then((data: string) => {
            console.log(data)
            state.result = data;
        })
}

export interface Attrs {
    name: string
}

class AppComponent implements ClassComponent<Attrs> {
    view({attrs: {}}: CVnode<Attrs>) {
     return m("main", [
         m("h2", {class: "title"}, "Enter your code"),
         m('select', {
             onchange: m.withAttr("selectedIndex", state.setIndex)
         }, targets.map((item: string, index: number) => {
             return m('option', {selected: state.index === index}, item);
         })),
         m('textarea[rows="20"]', {
             value: state.src,
             oninput: m.withAttr('value', (value: string) => {
                 state.src = value;
             }),
         }),
         m("button", {
             onclick: () => {
                 console.log(state);
                 render(state.target, state.src);
             },
         }, "render"),
         m("hr"),
         m("div.highlight", m.trust(state.result)),
     ])
    }
}

m.mount(document.getElementById("body"), AppComponent);
