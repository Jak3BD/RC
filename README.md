# RC
This is a simple, quick and dirty cli tool to create react components

Tested on Mac (arm64), most likely same on amd64 and linux amd64/arm64


## Installation
__Unix__
* Either download the compiled binary for your platform or download the project and compile it yourself
* `mkdir ~/bin`
* in my case i use the zsh shell `nano ~/.zshrc`
* add at the end `alias rc="~/bin/rc_<VERSION>_<OS>_<ARCH>"`
* `source ~/.zshrc`


## Usage

```bash
rc HelloWorld
rc myComponent/HelloWorld
```


## Options

### indent
* tabs / spaces
* count (number of spaces or tabs)

### quotes
* double / single quotes

### css module import
* true / false
* css / scss / custom

### Export Type
```jsx
// Default Functional Export (ID: 1)
export default function test() {}

// Named Functional Export (ID: 2)
export function test() {}

// Exported Arrow Function (ID: 3)
export const test = () => {};

// Default Exported Function (ID: 4)
function test() {}
export default test;

// Named Exported Function (ID: 5)
function test() {}
export { test };
```


### Example config (rc.json)
```json
{
    "indent": {
        "type": "tabs",
        "count": 1
    },
    "quotes": "double",
    "exportType": 1,
    "style": {
        "enable": true,
        "ext": "scss"
    }
}
```
