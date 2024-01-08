//lib call
import { insertHTML } from "https://cdn.jsdelivr.net/gh/jscroot/api@0.0.4/croot.js";
//internal call
import { url,id } from "../url/config.js";
import { getContentURL,getURLContentJS } from "../url/content.js";


export function runAfterHashChange(evt){
    insertHTML(getContentURL(),id.content,runAfterContent);
}

export function runAfterHeader(){
    insertHTML(url.template.navbar,id.navbar,runAfterNavbar);
}

async function runAfterNavbar(){
    let module = await import(url.view.navbar);
    module.main();
}

export async function runAfterContent(){
    let urljs = getURLContentJS();
    let module = await import(urljs);
    module.main();
    console.log(urljs);
}