



Group Members:
- Hiwot Derese
- Sosna Worku
- Hayat Tofik
- Tseganesh Yifru
- Ekram Kedir


# Project Description
##   Bencode

```
is a format in which the torrent file is written.To get useful parameters that can be used by the trackers and the other peers we need to apply different operations like parsing and marshaling.So under the bencode folder we have added he implementation for the parser and marshal and he corrponding test files for them.
- Parser.go - This file implements functionalities that are needed by the parser such as gettin the important informations from the .torrent file.After                                parsing the whole file , we can spot the URL of the tracker, the creation date (as a Unix timestamp), the name and size of the file, and a                                big binary blob containing the SHA-1 hashes of each piece, which are equally-sized parts of the file we want to download. The exact size of                              a piece varies between torrents, but they are usually somewhere between 256KB and 1MB. This means that a large file might be made up of                                  thousands of pieces. We’ll download these pieces from our peers, together, and boom, we’ve got a file!
```

## Usage

```js
import {subscribe} from '@github/paste-markdown'

// Subscribe the behavior to the textarea.
subscribe(document.querySelector('textarea[data-paste-markdown]'))
```

Using a library like [selector-observer][so], the behavior can automatically
be applied to any element matching a selector.

[so]: https://github.com/josh/selector-observer

```js
import {observe} from 'selector-observer'
import {subscribe} from '@github/paste-markdown'

// Subscribe the behavior to all matching textareas.
observe('textarea[data-paste-markdown]', {subscribe})
```

### Excluding `<table>`s

Some `<table>`s are not meant to be pasted as markdown; for example, a file content table with line numbers in a column. Use `data-paste-markdown-skip` to prevent it.

```html
<table data-paste-markdown-skip>
  ...
</table>
```

### Granular control for pasting as plain text

If you're wanting more granular support of pasting certain items as plain text by default, you can pass in the controls config at the `subscribe` level.

Our config support looks as follows:

```js
import {subscribe} from '@github/paste-markdown'

// Subscribe the behavior to the textarea with pasting URL links as plain text by default.
subscribe(document.querySelector('textarea[data-paste-markdown]'), {defaultPlainTextPaste: {urlLinks: true}})
```

In this scenario above, pasting a URL over selected text will paste as plain text by default, but pasting a table will still paste as markdown by default.

Only the `urlLinks` param is currently supported.

If there is no config passed in, or attributes missing, this will always default to `false`, being the existing behavior.

## Development

```
npm install
npm test
```

## License

Distributed under the MIT license. See LICENSE for details.




