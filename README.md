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

## Parameterübergabe

In Go werden Parameter by Value übergeben, d. h. es wird eine Kopie des Wert des Parameters übergeben. Das bedeutet, dass Änderungen am Wert eines Parameters in einer Funktion keine Auswirkung auf den Wert außerhalb der Funktion haben. Von dieser Regel gibt es nur wenige Ausnahmen - diese haben wir aber noch nicht kennengelernt.

## Rückgabewerte von Funktionen ignorieren

Der Aufrufer kann den Rückgabewert einer Funktion ignorieren (wenn er ihn nicht braucht), dazu verwendet man `_` anstatt des Variablennamen, wo man das Ergebnis der Funktion sonst speichern würde. Der `_` heißt in dem Zusammenhang **blank identifier**. Der \_ ist dabei nicht eine Konvention, sondern wirklich ein Merkmal von Go, d. h. man kann dann wirklich nicht mehr auf den Wert zurückgreifen.

Der Go Compiler wirft auch einen Fehler, wenn man eine ungenutzte Variable hat, daher ist es "oft" notwendig den \_ zu verwenden.

Und eine Funktion kann auch mehr wie einen Wert zurückliefern, dann kann/muss man auch alle "entgegennehmen" (und "notfalls" im \_ entsorgen). Um in der Deklaration der Funktion bekannt zu geben, dass sie mehrere Rückgabewerte hat, klammert man die Rückgabewerte ein: `func mehrere(tier string) (string,int,int) { ... return 'hallo',1,2}`.

## Rückgabewerte von Funktionen benennen

Man kann Rückgabewerte auch benennen. Das ist sinnvoll um die Lesbarkeit (insbesondere bei längeren Funktionen) zu erhöhen.

```golang
func getCoords() (x, y int){
  // x and y are initialized with zero values

  return // automatically returns x and y
}
```

Das return in der obigen Form wird als **naked return** (oder **implicit return**) bezeichnet (weil man die Namen der Rückgabewerte nicht extra angibt). Das sollte nur in kurzen Funktionen verwendet werden, so man die Funktionsdeklaration beim return noch "in Sicht" hat.

Und das entspricht dieser "Langform" (die trotzdem weniger Übersichtlich ist) - daher sollte man die vorstehende Variante bevorzugen:

```golang
func getCoords() (int, int){
  var x int
  var y int
  return x, y
}
```

## Early returns

Sobald das Programm bei der Ausführung auf ein return trifft, wird die aktuelle Funktion verlassen. Wenn danach noch Code in der Funktion kommt, wird das als **early return** bezeichnet. Das macht Sinn, wenn das early return von einer Bedingung abhängig ist - also nicht immer "schlagend" wird. In so einem Fall kann es helfen, die Bedingungen einfacher zu gestalten - weil man weiß dass die Bedingung, die das early return ausgelöst hätte, danach nicht mehr erfüllt sein kann (und man das an der Stelle dann nicht mehr prüfen muss). Das early return wird auch als **guard clause** bezeichnet.

## Funktionen als Typen

In Go sind Funktionen auch gleichzeitig Typen (wie `string` oder `int`). Man kann sie daher überall dort verwenden, wo man Typen verwenden kann.

```golang
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
  firstResult := arithmetic(a, b)
  secondResult := arithmetic(firstResult, c)
  return secondResult
}
```

Die Funktion aggregate erwartet drei `int` und dann eine Funktion arithmetic (die selbst zwei int als Paramter erwartet und einen int als Rückgabe liefert). Die Funktion aggregate wendet die Funktion arithmetic zuerst auf ihre ersten beiden Parameter an und anschließend die Funktion arithmetic ein weiteres mal auf das erste Zwischenergebnis und auf ihren dritten Parameter. Das Ergebnis dieses zweiten Aufrufs wird dann das Gesamtergebnis von aggregate. So ungefähr kann man sich die reduce Methode in Javascript für Arrays vorstellen (nur ist die universeller - was aber daran liegt, das Javascript keine statischen Typen hat).

## Anonyme Funktionen

Dabei handelt es sich um Funktionen, denen man keinen Namen gegeben hat. Das passiert oft an Stellen, wo eine Funktion eine andere Funktion als Parameter erwartet (so wie auch in Javascript).

## defer

Das ist ein "Alleinstellungsmerkmal" von Go. Mit `defer` kann man eine Funktion definieren, die aufgerufen wird unmittelbar bevor die umgebende Funktion beendet wird. Wenn in einer Funktion mehrere `defer` Aufrufe existieren, werden sie in **umgekehrter** Reihenfolge am Ende der Funktion ausgeführt (d. h. das zuletzt erreichte defer wird als erstes ausgeführt und das erste zuletzt). Man verwendet das im Normalfall um Resourcen wieder freizugeben. Man muss sich aber in Go nicht extra darum kümmern, sich zu merken was man aufräumen muss, sondern man erstellt das "Aufräumen" gleich nach dem Anlegen der Ressource - und Go kümmert sich dann darum es zur richtigen Zeit auszuführen. Es ist "ähnlich" wie finally in einem try/catch Block, das auch immer ausgeführt wird, egal wie der try/catch Block beendet wird.

## Block scope

Go ist eine block scoped Programmiersprache, das bedeutet, dass eine Variable nur innerhalb ihres Blocks gültig ist - wie auch Javascript. Blöcke werden dabei mittels den `{}` definiert. Und man kann neben den "normalen" Blöcken (z. b. bei Funktionen oder if) auch eigene Blöcke definieren. Dies dient im Normalfall nur dazu Variablen weiter zu kapseln.

## Closure

Ein Closure ist eine Funktion, die auf eine Variable von außerhalb ihres Funktionsblocks zugreift. Die Funktion kann dabei sowohl lesend als auch schreibend auf diese Variable zugreifen. Im untenstehenden Beispiel wird für jeden Aufruf von concatter der übergebene String an die bisher übergebenen Strings angehängt:

```golang
func concatter() func (string) string {
  doc :=""
  return func (word string) string {
    doc += word + " "
    return doc
  }
}
```

So kann man das **NICHT** verwenden (weil concatter ja eine Funktion zurückgibt und das irgendwo "gespeichert" werden muss).

```golang
func main() {
  concatter("Hallo")
  concatter("Welt")
  concatter("von")
  fmt.Println(concatter("Powidl")) // Ergibt die Ausgabe "Hallo Welt von Powidl"
}
```

Statt dessen verwendet man es so:

```golang
func main() {
	f := concatter()
	f("Hallo")
	f("Welt")
	f("von")
	fmt.Println(f("Powidl")) // Ergibt die Ausgabe "Hallo Welt von Powidl"
}
```

Wichtig dabei ist, dass bei der Initialisierung von f **KEIN** Parameter an concatter übergeben wird - weil ja concatter keinen Parameter erwartet, sondern eine Funktion zurückliefert, die einen Parameter erwartet.

```golang
func concatter2(init string) func (string) string {
  doc := init + " "
  return func (word string) string {
    doc += word + " "
    return doc
  }
}
```

Jetzt erwartet concatter2 auch einen Parameter (und das wird als "Startwert" für den zu bildenden string verwendet).

In dem Fall muss man main() so schreiben:

```golang
func main() {
  f := concatter2("Hallo")
  f("Welt")
  f("von")
  fmt.Println(f("Powidl")) // Ergibt ebenfalls "Hallo Welt von Powidl"
}
```

# Kapitel 4: structs

`structs` werden zur Repräsentation von strukturierten Daten verwendet.

## Definition von structs

```golang
type car struct {
  brand string
  model string
  doors int
  mileage int
}
```

Die Definition wird mit dem Schlüsselwort `type` eingeleitet, dann folgt der Name des structs gefolgt von `struct` (`struct`) ist also selbst ein Datentyp. Dann gefolgt in geschwungenen Klammern die einzelnen Elemente, die das struct umfasst.

Und weil struct selbst ein Datentyp ist, kann man sie auch ineinander verschachteln.

```golang
type car struct {
  brand string
  model string
  doors int
  mileage int
  frontWheel wheel
  backWheel wheel
}

type wheel struct {
  radius int
  material string
}
```

## Erzeugen von Variablen mit einem bestimmten struct Typ

Entweder verwendet man `var audi car` um die Variable audi vom Typ car zu definieren - oder in einer Funktion `audi := car{}`.

## Zugriff auf Elemente des struct

Auf die einzelnen Elemente des struct wird dann mittels . zugegriffen, z. b. `audi.brand = "Audi"`.

## Anonyme structs

Man kann ein struct auch definieren ohne dass man ihm einen Namen gibt, dann kann man diesen struct Typ an keiner anderen Stelle ansprechen. Um diesen struct mit Werten zu befüllen, kann man eine zweite geschwungene Klammer (nach der Festlegung der Typen der Elemente) mit den jeweiligen Werten der Elemente für diese konkrete Variable verwenden:

```golang
myCar := struct {
  brand string
  model string
} {
  brand: "Toyota",
  model: "Camry",
}
```

**Wichtig**: Nach dem letzten Element in der "Zuweisungs-{}" muss ebenfalls ein `,` stehen (was in der Powershell einen Fehler auslösen würde, bei Javascript erlaubt ist, ist bei Go verpflichtend - es kommt ein Compilerfehler, wenn das nicht gemacht wird).

Diese Form der Wertinitialisierung kann man genauso bei benannten structs verwenden, dann sind die leeren {} eben nicht leer sondern darin belegt man die einzelnen Elemente mit den Startwerten.

Generell sind benannte structs zu bevorzugen, es kann aber Sinn machen anonyme structs zu verwenden - wenn man zum Ausdruck bringen will, dass dieser struct an keiner anderen Stelle verwendet werden soll.

## embedded structs

Go ist keine objektorientierte Sprache. Mit embedded structs kann man aber zumindest Attribute zwischen verschiedenen structs teilen.

```golang
type truck struct {
  car
  bedSize int
}
```

Damit "erbt" das struct truck alle Elemente von car und hat zusätzlich noch ein Element bedSize. Das ist etwas anderes wie wheel (innerhalb von car). car hat ein bzw. zwei wheels (frontWheel und backWheel), während `truck` ein `car` **ist**. In objektorientierten Sprachen spricht man auch von **has a** (wheel) und **is a** (truck).

Wobei man bei den geerbeten Elementen zwischen der Initialisierung und dem Zugriff unterscheiden muss. Bei der Initialisierung muss man über das "Basiselement" gehen, beim Zugriff kann man auf die "geerbten" Elemente ohne "Zwischenschritt" zugreifen.

Die Initialisierung von einem truck könnte so aussehen

```golang
m18:=truck{
  car: car{
    brand:"Steyr",
    model:"M18",
    doors:4,
    mileage:0,
    frontWheel:wheel{
      radius:35,
      material:"gum",
    },
  },
  bedSize:0,
}
```

Und der Zugriff dann so: `fmt.Printf("Der Truck hat eine Bedsize von %v und hat %v Türen",m18.bedSize,m18.doors).`. Der Zugriff auf doors (aus car) ist auch über `m18.car.doors` möglich (das kann interessant sein, wenn man in Truck direkt auch eine Eigenschaft doors definiert - wobei man sich solche Dinge sehr genau überlegen sollte, weil es den Code nicht gerade "verständlich" macht).

## Methoden

Obwohl Go nicht objektorientiert ist, hat es Methoden. Methoden sind spezielle Funktionen (mit einem zusätzlichen "Parameter" - dem `receiver`). Der `receiver` wird bei der Deklaration vor dem Funktionsnamen (bzw. nach dem Schlüsselwort func geschrieben) - und zwar ebenfalls in runden Klammern Name Typ:

Im folgenden Beispiel wird die Methode area für den Typ rect erstellt:

```golang
type rect struct {
  width int
  height int
}

func (r rect) area() int {
  return r.width*r.height
}

var r=rect{
  width:3,
  height:7,
}
func main() {
  fmt.Printf("Die Fläche des rect %v x %v ist %v\n",r.width,r.height,r.area()); // Ergibt "Die Fläche des rect 3 x 7 ist 21"
}
```

## Memory Layout

Ein struct liegt als zusammenhängender Block im Speicher. Man nutzt den Speicher "optimal" wenn man die Elemente eines structs nach ihrer Speichergröße deklariert (vom größten zum kleinsten). Im Normalfall muss man sich aber nicht allzu viele Gedanken darüber machen, aber bei großen Arrays mit vielen Elementen, die alle structs sind, kann es einen nennenswerten Einfluss auf den Speicherbedarf und auch auf die Ausführungsgeschwindigkeit haben. Wobei mittlerweile der Go Compiler intelligenter zu sein scheint, aber man kann mit der unterschiedlichen Reihenfolge immer noch "deutliche" Unterschiede im Platzbedarf erzeugen.

Eine Möglichkeit den Platzbedarf eines Objekts im Speicher zu messen bietet das Package `reflect`: `reflect.Type(truck{}).Size()` - liefert die Größe eines leeren truck structs im Speicher.

## Empty struct

Ein empty struct ist der kleinstmögliche Typ in Go. Er belegt 0 Bytes Memory. Er wird wie folgt definiert (einmal als anonymer struct und einmal als Typ emptyStruct):`

```golang
// anonymous empty struct type
empty := struct{}{}

// named empty struct type
type emptyStruct struct{}
empty := emptyStruct{}
```

# Kapitel 5: interfaces

Interfaces erlauben es sich darauf zu konzentrieren, welche Methoden ein Typ zur Verfügung stellt (und nicht wie er genau gebaut ist). In Go bestehen interfaces nur aus Methoden, nicht auch aus Attributen. Und ein Typ implementiert ein interface, wenn er alle Methoden des Interfaces mit den richtigen Parameter- und Rückgabetypen hat. Ein Interface in Go ist daher eine Sammlung von Methoden. Und ein Typ kann beliebig viele Interfaces "gleichzeitig" implementieren.

Der zero-value von interfaces ist `nil` (und weil `error` in Go ein Interface ist, checkt man auf das "Nichtvorliegen" eines Fehlers indem man den error mit `nil` vergleicht).

## Definition

```golang
type interfacename interface {
  methode(parameterliste) returntyp
  methode(parameterliste) returntyp
}
```

## Implementierung

Interfaces werden nicht explizit implementiert (d. h. man gibt bei einem Typ nicht an, dass er ein interface implementiert). Das passiert automatisch - wenn ein Typ alle Voraussetzungen eines Interfaces erfüllt hat er dieses Interface implementiert.

## Verwendung

Der Interfacetype kann dann als ganz normaler Type verwendet werden (beispielsweise in der Parameter- oder Rückgabedefinition von Funktionen).

Und es gibt auch - wie bei `struct` ein **empty interface**.

## Named Parameters

Bei interfaces ist es noch wichtiger, dass man Parametern (und die return values) mit Namen versieht, damit jemand, der ein Interface implementiert weiß, wofür die jeweiligen Dinge gedacht sind.

## Type assertions

Manchmal will/muss man (bei interfaces) auf den "zugrundeliegenden" Typ zugreifen. Dazu kann man Type assertions verwenden.

```golang
type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

c, ok := s.(circle)
if !ok {
	// log an error if s isn't a circle
	log.Fatal("s is not a circle")
}

radius := c.radius
```

Die Type assertion im obigen Code ist das `.(circle)`. Man kann sich die Type assertion als "spezielle" Methode vorstellen (die keinen Namen hat) und als Parameter den Typ hat, auf den geprüft werden soll. Sie hat zwei Rückgaben. Als erstes das gecastete Objekt, als zweites ein boolean, dass angibt ob das casting erfolgreich war. Im Beispiel oben ist c am Ende ein circle, wenn s ein circle war. **ACHTUNG**: Wenn s kein circle (sondern ein anderes shape) war, ist c trotzdem vom Typ circle! D. h. man muss selbst auf den zweiten Parameter prüfen und dementsprechend richtig reagieren!

## Type switches

Type switches sind normale switch Statements, aber der Fall wird aufgrund von Typen unterschieden.

```golang
func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}

func main() {
	printNumericValue(1)
	// prints "int"

	printNumericValue("1")
	// prints "string"

	printNumericValue(struct{}{})
	// prints "struct {}"
}
```

Wenn man den Typ selbst nicht als Variable braucht (Wir brauchen ihn im obigen Beispiel, damit wir ihn ausgeben können), kann man die erste Zeile auch so schreiben `switch num.(type) {`.

## Clean interfaces

Es ist wichtig seine Interface klar (und übersichtlich) zu halten. Das kann sehr schnell sehr schwer werden. Hier sind ein paar Grundregeln, denen man nach Möglichkeit folgen sollte:

1. **Interfaces sollen so klein wie möglich sein**: Sie sollen das minimal notwendige Verhalten darstellen, dass notwendig ist um eine Idee oder ein Konzept zu repräsentieren.
2. **Interfaces sollen kein Wissen über die implementierenden Typen benötigen**: Sie stellen dar, was jemand mitbringen muss, um als Typ des Interfaces zu gelten. Sie sollen über diesen Jemand nichts wissen müssen.
3. **Interfaces sind keine Klassen**: Sie haben keine Konstruktoren oder Destructoren, sie behandeln keine Daten (sie schreiben nur vor, wie man mit etwas, dass das Interface implementiert, kommunizieren kann bzw. muss). Sie kümmern sich auch nur um das "was" und nicht um das "wie".
