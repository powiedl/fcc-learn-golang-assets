# Assets for "Learn Go" on FreeCodeCamp

This is a snapshot of the code samples for the ["Learn Go" course](https://boot.dev/courses/learn-golang) on [Boot.dev](https://boot.dev) at the time the video for FreeCodeCamp was released on YouTube. If you want the most up-to-date version of the code, please visit the official [Boot.dev course](https://boot.dev/courses/learn-golang). Otherwise, if you're looking for the files used in the video, you're in the right place!

- [Course code samples](/course)
- [Project steps](/project)

## License

You are free to use this content and code for personal education purposes. However, you are _not_ authorized to publish this content or code elsewhere, whether for commercial purposes or not.

# Meine Mitschrift

Hier schreibe ich mit, was ich für notierwürdig halte. Wahrscheinlich kommen mir am Ende viele der Notizen als "überflüssig vor". Ich werde auch versuchen, die einzelnen Notizen "kurz" zu halten. Ich will sie am Ende in mein noch zu erstellendes "Kurs Notizen" Programm importieren (die Idee zu diesem Programm ist bei der Verwendung von udemy entstanden - dort darf eine Notiz nur 1000 Zeichen lang sein - und das kann manchmal zu kurz sein).

# Kapitel 1: Variablen

## Variablen Deklarieren und initialisieren

Variablen müssen in Go einen Typ haben. Ein paar "Basis"typen sind `byte`, `bool`, `int`, `float64` und `string`. Daneben gibt es für diese Typen auch "Spezialfälle" (diese unterscheiden sich in der Speichermenge und der Interpretation des Speichers). Solange es möglich ist, **sollte man den jeweiligen Basistypen (int, uint, float64 und complex128 - für Zahlen) verwenden**, da es sonst relativ schnell relativ aufwändig werden kann (man muss gegebenenfalls die Variablen von einem Subtyp in einen anderen konvertieren).

`strings` müssen mit `"` umschlossen werden - `'` sind nicht erlaubt.

Beispiele für Subtypen sind `uint16` (positive ganzzahlige Zahlen mit 16bit Länge), `float32` (Fließkommazahlen mit 32bit Länge).

Variablen werden mittels des Schlüsselworts `var` definiert. Dann folgt der Name und dann der Typ: `var name string` oder `var kosten float64`.

Wenn man den Variablen dann einen Wert zuweisen will macht man das in einer extra Zeile `name="powidl"` oder `kosten = 123.4`.

Innerhalb von Funktionen kann man auch den walrus Operator `:=` (**short assignment operator**) verwenden. Dieser deklariert eine Variable und weist ihr auch gleich einen Wert zu (den Typ der Variablen bestimmt Go anhand des Werts) `vorname := "pow"` (vorname ist vom Typ string) oder `muede := true` (muede ist vom Typ bool).

Man kann auch mehrere Variablen in einer Zeile definieren und ihnen einen Wert zuweisen. Die Typen können ebenfalls unterschiedlich sein. Dazu trennt man vor dem `:=` die Variablennamen mit `,` und danach die Werte ebenfalls mit `,`: `averageOpenRate, displayMessage := 23,"is the average open rate of your messages"`

## Konstanten

Konstanten werden mit dem Schlüsselwort `const` definiert. Bei ihnen kann man `:=` **nicht** verwenden, aber ihnen muss in derselben Zeile ein Wert zugewiesen werden `const pi = 3.14159`.

## Kommentare

Einzeilige Kommentare beginnen mit `//`, mehrzeilige Kommentare beginnen mit `/*` und enden mit `*/`. Kommentare werden vom Go Compiler komplett ignoriert. Konstanten müssen zur Entwicklungszeit bekannt sein (sie können nicht erst zur Laufzeit des Programms einen Wert annehmen).

## Generelle Struktur eines Go Programms

Wenn ein Programm mit `package main` beginnt, handelt es sich bei der Datei um ein eigenständiges Programm (bzw. in einem komplexeren Projekt um die Hauptdatei).

Dann folgen die Importe der Bibliotheken, die man in seinem Programm benutzt (z. b. `fmt` für Ausgaben).

Und dann kommt `func main() { ... }` - der Einsprungspunkt für ein Go Programm.

## Go Programme kompilieren

Go Programme müssen kompiliert werden. Dazu verwendet man den Befehl `go`: `go build main.go`.

## Go ist statisch typisiert

Das bedeutet, dass der Typ von Variablen bzw. Werten zur Entwicklungszeit feststeht. Das hat den Vorteil, dass der Editor während der Entwicklung auf Typfehler prüfen kann und diese nicht erst zur Laufzeit auffallen. Der "Nachteil" ist, dass es "mühsamer" wird Programme zu schreiben, weil man nicht zwischen verschiedenen Typen für ein und dieselbe Variable wechseln kann (aber dieses "mühsame" zahlt sich im Normalfall recht schnell aus!).

## Verbinden von Strings

Wie in vielen anderen Programmiersprachen auch, kann man in Go strings mit dem `+` aneinanderfügen. Aber man kann nicht einen string und eine Zahl addieren.

## Formattieren von Strings

Go folgt der Formattierung von strings von `C`. Es gibt `fmt.Printf` (Ausgabe des Strings auf Standard Out) und `fmt.Sprintf()` (liefert einen formatierten String zurück - den man dann gegebenenfalls selbst ausgeben muss).

Die Syntax von beiden Befehlen ist gleich. Als erster Parameter wird ein Formatierungsstring übergeben und danach die einzelnen Werte. Im Formatierungsstring kann man Formatierungsoptionen angeben. Diese beginnen mit einem % gefolgt von einem oder mehreren Zeichen:

- `%v` der Wert wird in seiner Representanz in Go ausgegeben (in diesem Fall muss man nicht wissen, ob es sich bei dem Wert um einen String oder eine Zahl oder ... handelt)
- `%s` der Wert muss ein `string` ausgegeben
- `%d` der Wert muss ein `int` sein
- `%t` der Wert muss ein `bool` sein (und es wird `true` oder `false` ausgegeben)
- `%f` der Wert muss ein `float`sein. Hier gibt es die Möglichkeit zwischen % und f weitere Merkmale für die Ausgabe festzulegen (%.2f legt fest, dass die Ausgabe genau zwei Nachkommastellen haben soll - es wird gerundet).
- `%T` der Typ des Werts wird ausgegeben (anstatt des Werts, z. b. `"%T",4` ... ergibt `int`).

## rune

Ein `rune` ist ein `int32`, aber es wird als ein (UTF8)-Zeichen interpretiert. Go verwendet standardmäßig rune für Strings. Wenn man die Länge eines rune mit `len()` abfrägt, erhält man 4 als Ergebnis. wenn man aber die Funktion `utf8.RuneCountInString()` verwendet erhält man 1 als Ergebnis.

# Kapitel 2: Conditionals

## if

Die Bedingung in einer `if` Anweisung wird **nicht** in Klammern geschrieben. Es gibt kein `then`, aber `else if` und `else`. Und es kann auch mehrere `else if` in einer `if` Anweisung geben. Sobald eine der Bedinungen erfüllt ist, werden die anderen nicht mehr ausgewertet.

Eine `if` Anweisung kann auch ein initiales Statement enthalten (mit ; vor der Bedingung). Variablen, die in dem initialen Statement definiert werden gelten nur innerhalb des `if`s. Die Vorteile dieses initialen Statements sind, dass der Code etwas kürzer ist und - vor allem - dass die Variable nur in dem `if` gilt. Das folgt dem Prinzip, dass man Variablen immer so lokal wie möglich definieren sollte - um die Gefahr von Seiteneffekten zu reduzieren.

## Vergleichsoperatoren in Go

- `==` gleich
- `!=` ungleich
- `<` kleiner
- `>` größer
- `<=` kleiner gleich
- `>=` größer gleich

## switch

Mit einem `switch` kann man das gleiche wie mit mehreren if - else if - else if - else erreichen. Im Normalfall ist das switch aber besser lesbar.

```golang
var creator string
switch os {
  case "linux":
    creator = "Linus Torvalds"
  case "windows":
    creator = "Bill Gates"
  default:
    creator = "Unknown"
}
```

Im Gegensatz zu Javascript wird in Go nur ein `case` ausgeführt, es wird also nicht zum nächsten fortgefahren. Will man doch, dass mit dem nächsten case weitergemacht werden soll, muss man `fallthrough` verwenden (das aber wieder nur bis zum dann nächsten case gilt).

Bei `switch` wird für den "Standardfall" nicht `else` sondern `default` verwendet.

# Kapitel 3: Funktionen

Funktionen werden mit dem Schlüsselwort func eingeleitet. Dann folgen die Parameter in `()`. Eine Funktion kann null oder mehrere Parameter haben. Die () müssen immer angegeben werden - auch wenn eine Funktion keine Parameter hat.

Die Parameter werden durch `,` getrennt. Zuerst kommt der Parametername und dann der Typ. Wenn mehrere Parameter den gleichen Typ haben, reicht es diesen einmalig nach dem letzten gleichartigen Parameter anzugeben `func add(a,b int) int { return a+b }`.

Nach dem Parameterblock kommt der Rückgabetyp der Funktion (welchen Typ gibt die Funktion zurück).

Daran anschließend kommt der eigentliche Funktionsblock (was soll gemacht werden, wenn die Funktion aufgerufen wird) wird mit {} umschlossen.

## Deklarationen

Im Gegensatz zu C wird in Go der Typ nachgestellt. Das entspricht eher dem Sprachgebrauch in Englisch "variable x of type int" `var x int`.
