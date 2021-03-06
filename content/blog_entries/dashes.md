# Dashes — You should learn to use them.

I guess this will be the first entry of what I'm going to call my blog. Or a weblog, if you prefer.

I have been wanting to start a blog for a long time. I just never had a place (nor time, which I currently actually don't have, but shhh) to keep it. I guess this is it.

Anyway, what pushed me over the edge was this: I was browsing StackExchange while I was talking to someone about my `vim` configuration. I had `:set list` on while sending the screenshot, and well. They asked about my "underscores". I use `␣` (U+2423) for visual spaces.

I tripped over an [answer](https://ux.stackexchange.com/a/91282/132106) which I had to comment on. They called the hyphens in the url for DASHES. Who... in... their... right... mind!? Hyphens are NOT dashes!

Apparently, there are a bunch of well defined dashes.

`xpath -q -e "//char[@Dash='Y']/@*[name()='cp' or name()='na']" ucd.all.flat.xml | sed -E 's/"|^ |na=//g' | sed 's/cp=/U+/g' | paste -d ':\t' - /dev/null -`

```
U+002D:	HYPHEN-MINUS
U+058A:	ARMENIAN HYPHEN
U+05BE:	HEBREW PUNCTUATION MAQAF
U+1400:	CANADIAN SYLLABICS HYPHEN
U+1806:	MONGOLIAN TODO SOFT HYPHEN
U+2010:	HYPHEN
U+2011:	NON-BREAKING HYPHEN
U+2012:	FIGURE DASH
U+2013:	EN DASH
U+2014:	EM DASH
U+2015:	HORIZONTAL BAR
U+2053:	SWUNG DASH
U+207B:	SUPERSCRIPT MINUS
U+208B:	SUBSCRIPT MINUS
U+2212:	MINUS SIGN
U+2E17:	DOUBLE OBLIQUE HYPHEN
U+2E1A:	HYPHEN WITH DIAERESIS
U+301C:	WAVE DASH
U+3030:	WAVY DASH
U+30A0:	KATAKANA-HIRAGANA DOUBLE HYPHEN
U+FE31:	PRESENTATION FORM FOR VERTICAL EM DASH
U+FE32:	PRESENTATION FORM FOR VERTICAL EN DASH
U+FE58:	SMALL EM DASH
U+FE63:	SMALL HYPHEN-MINUS
U+FF0D:	FULLWIDTH HYPHEN-MINUS
```

I... am... disgusted. Hyphens are dephined as dashes in Unicode.
Nevertheless, I decided to look at the differences between dashes and hyphens according to the Unicode Consortium. There are no dashes in the Unihan dataset, unsurprisingly, but there are 25 of them in the standard set. Same applies for the following on hyphens, except we have 28 in total.

`xpath -q -e "//char[@Hyphen='Y' or @Dash='Y']/@*[name()='cp' or name()='na' or name()='Hyphen' or name()='Dash']" ucd.all.flat.xml | sed -E 's/"|^ |na=//g' | sed 's/cp=/U+/g' | sed -z 's/\n/:/g' | sed 's/U+/\nU+/g' | sed -E 's/Dash=N|Hyphen=N//g' | column -t -s:`

Call me out on excessively long Bash commands. You'll get used to it.

```
U+002D  HYPHEN-MINUS                            Dash=Y  Hyphen=Y
U+00AD  SOFT HYPHEN                                     Hyphen=Y
U+058A  ARMENIAN HYPHEN                         Dash=Y  Hyphen=Y
U+05BE  HEBREW PUNCTUATION MAQAF                Dash=Y
U+1400  CANADIAN SYLLABICS HYPHEN               Dash=Y
U+1806  MONGOLIAN TODO SOFT HYPHEN              Dash=Y  Hyphen=Y
U+2010  HYPHEN                                  Dash=Y  Hyphen=Y
U+2011  NON-BREAKING HYPHEN                     Dash=Y  Hyphen=Y
U+2012  FIGURE DASH                             Dash=Y
U+2013  EN DASH                                 Dash=Y
U+2014  EM DASH                                 Dash=Y
U+2015  HORIZONTAL BAR                          Dash=Y
U+2053  SWUNG DASH                              Dash=Y
U+207B  SUPERSCRIPT MINUS                       Dash=Y
U+208B  SUBSCRIPT MINUS                         Dash=Y
U+2212  MINUS SIGN                              Dash=Y
U+2E17  DOUBLE OBLIQUE HYPHEN                   Dash=Y  Hyphen=Y
U+2E1A  HYPHEN WITH DIAERESIS                   Dash=Y
U+301C  WAVE DASH                               Dash=Y
U+3030  WAVY DASH                               Dash=Y
U+30A0  KATAKANA-HIRAGANA DOUBLE HYPHEN         Dash=Y
U+30FB  KATAKANA MIDDLE DOT                             Hyphen=Y
U+FE31  PRESENTATION FORM FOR VERTICAL EM DASH  Dash=Y
U+FE32  PRESENTATION FORM FOR VERTICAL EN DASH  Dash=Y
U+FE58  SMALL EM DASH                           Dash=Y
U+FE63  SMALL HYPHEN-MINUS                      Dash=Y  Hyphen=Y
U+FF0D  FULLWIDTH HYPHEN-MINUS                  Dash=Y  Hyphen=Y
U+FF65  HALFWIDTH KATAKANA MIDDLE DOT                   Hyphen=Y
```

Another small detail is that `Dash` or `Hyphen` are only defined as `"N"` for the group E000–F8FF (I used an U+2013 here!) in the grouped version, but are well defined for 107 356 elements in the full and flat version. There are a total of 107 364 defined characters. Properly, 8 of them are defined true for both.

Another detail, instead of a space (U+0020), a figure space (U+2007) was used, as a thousands separator. The International System of Units has declared a space to be used as a decimal separator, and while [Wikipedia — Thin Space](https://en.wikipedia.org/wiki/Thin_space) says they use a thin space to denote this, I disagree. I believe a *figure space* is the correct option.

A shorter, more useful list:

| Unicode | Symbol   | Usage |
| :------ | :------: | ----- |
| U+002D  | -        | It's what you get when hitting the numpad! **A compromise** needed during the 7-bit era. |
| U+00AD  | ­         | You're probably not even seing this one, as it's only used to **allow renderers to break a line**, with a true hyphen instead! |
| U+2010  | ‐         | **Hyphens.** Used to split super-words (like this one!) into different parts. Also used at the end of a line to mark the splitting of a word. |
| U+2011  | ‑         | Visibly, the same as above. Except the renderer **mustn't add a newline** after it. |
| U+2012  | ‒         | **Phone numbers!** Used mostly in USA, though. (Witness how we both have figure dashes and figure spaces) |
| U+2013  | –         | I know this one by heart. En dash. Used to mark periods. Or intervals. **Anything with a range**, really. I've been living during the years 1997–2019. In latex you can use a double hyphen-minus.
| U+2014  | —         | This one is also a nice one to know! <br /> Used to add **comments** —much like parentheses—, to **halt mid-sente—**! Ah! I almost forgot: It's also used commonly in bibliographies (or other tabular data) to **repeat the previous column**'s data.<br /> So... many... usages. Also, look at the title of this blog entry. In latex it suffises to use three hyphen-minus symbols.
| U+2015  | ―         | Used at the beginning of a new speaker in dialogues.
| U+2053  | ⁓        | Used in examples of dictionaries as a placeholder for the word being defined. I might come back to variations of tildes sometime. |
| U+207B  | ⁻         | Very useful in mathematics. The ```arcsin(x)``` function is often represented as ```sin⁻ⁱ(x)```.  |
| U+208B  | ₋         | Similarly, useful if you need to type ```xₖ₋₁```. |
| U+2212  | −         | A true minus sign. ```x = A−B``` |

Oh, and don't forget: You can easily input Unicode characters in Ibus with <kbd>Ctrl</kbd><kbd>Shift</kbd><kbd>U</kbd>+<kbd>&lt;Unicode&gt;</kbd>+<kbd>␣</kbd>.

I have had some problems in the past with Qt5 not allowing unicode entry, but I usually fix this by installing the libraries for qt-ibus. On Manjaro the package is called `ibus-qt`, and needs this in my `.bashrc` file:
```bash
export GTK_IM_MODULE=ibus
export XMODIFIERS=@im=ibus
export QT_IM_MODULE=ibus
ibus-daemon -drx
```

Agh, it's 4:27 a.m. Time to sleep, I guess.
