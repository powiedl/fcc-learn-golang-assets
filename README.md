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

Außerdem ist Go stark typisiert, d. h. es finden keine automatischen Typumwandlungen durch den Compiler statt - man kann also nicht ein int und ein float addieren, sondern man muss explizit einen der beiden Werte in den Typ des anderen umwandeln.

Und auch wenn man konstante Werte angibt haben sie einen bestimmten Typ (**3** ist vom Typ `int`, **3.0** vom Typ `float64`). Eine "Ausnahme" ist **0**, das kann sowohl ein `int` als auch ein `float64` sein.

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

## logische Operatoren

- `&&` AND (beide Seiten müssen erfüllt sein, damit die gesamte Bedingung erfüllt ist)
- `||` OR (zumindest eine der beiden Seiten muss erfüllt sein, damit die gesamte Bedingung erfüllt ist)
- `!` NOT (negiert den Wert der Bedingung)

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

Wenn man beim Funktionsaufruf nach dem letzten Parameter vor der `)` eine neue Zeile anfängt (damit man die schließende Klammer "schön" formatieren kan, muss man nach dem letzten Parameter ein `,` verwenden).

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

## Higher order function

Darunter versteht man eine Funktion, die als Parameter eine Funktion entgegen nimmt. Im Normalfall wird diese Parameterfunktion, dann in der eigentlichen Funktion irgendwie aufgerufen.

Die "Parameter"funktion wird als **first-class Function** bezeichnet. Das heißt die Higher order function nimmt eine first-class function entgegen und führt sie zu einem späteren Zeitpunkt aus.

Die eigentliche Funktion muss keine Ahnung haben, was die übergebene Funktion macht. Die übergebene Funktion muss nur die richtige Signatur haben. Und es ist im Prinzip vergleichbar mit Callback-Funktionen in Javascript.

Ein häufiger Anwendungsfall sind jegliche Eventhandler. Da will man, dass der angegebene Code nicht sofort läuft, sondern wenn in der Zukunft irgendetwas bestimmtes passiert.

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

Und man muss auch nicht vor der Verwendung eine Variable/ein Objekt vom passenden Typ erzeugen, man kann das auch direkt beim Aufruf erledigen (man erzeugt ein anonymes Objekt von einem passenden Typ) - wenn man das nachher nicht noch woanders benötigt.

Im folgenden Beispiel nutzt man das beim Aufruf von printArea in main.

```golang
type shape interface {
  area() float64
}

type rect struct {
  width float64
  height float64
}
func (r rect) area() float64 {
  return r.height * r.width
}

func printArea(s shape) {
  fmt.Printf("The area of the shape is %v\n",s.area())
}

func main() {
  printArea(rect{
    width:12.4,
    height:4.3,
  })
}
```

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
3. **Interfaces sind keine Klassen**: Sie haben keine Konstruktoren oder Destructoren, sie behandeln keine Daten (sie schreiben nur vor, wie man mit etwas, dass das Interface implementiert, kommunizieren kann bzw. muss). Interfaces kümmern sich nicht um das erreichte Verhalten, sie legen nur fest, wie etwas aussehen muss, damit man es als "zu diesem Interface passend" betrachten kann.

# Kapitel 6: Errors

In Go werden Fehler als `error` Werte ausgedrückt. Ein `error` ist dabei alles, was das built-in `error` interface implementiert:

```golang
type error interface {
  Error() string
}
```

D. h. ein `error` muss nur eine Voraussetzung erfüllen. Er muss eine Methode Error haben, die keinen Parameter entgegennimmt und einen `string` zurückliefert.

Wenn etwas in einer Funktion fehlschlagen kann sollte sie als letzten Returnvalue immer ein error zurückliefern. Und der aufrufende Code sollte in so einem Fall immer prüfen, ob ein error oder `nil` (=kein Fehler aufgetreten) zurückgegeben wurde.

```golang
import "errors"

func doSomething(p int) (int,error) {
  if p % 2 == 1 { return 1,errors.New("p ist ungerade!") }
  return 0, nil
}
```

In dem Beispiel verwendet man einen Standardfehler aus dem Package errors. Er wird mit der Methode New erzeugt, welche einen Parameter als String - eben die Feherlmeldung - erwartet.

## eigene Fehler

Da `error` "nur" ein Interface ist kann man seine eigenen Fehlerobjekte erstellen (diese müssen nur eine Methode `Error() string` implementieren).

```golang
type userError struct {
    name string
}

func (e userError) Error() string {
    return fmt.Sprintf("%v has a problem with their account", e.name)
}

func sendSMS(msg, userName string) error {
    if !canSendToUser(userName) { // diese "Prüffunktion" ist in diesem Beispiel nicht implementiert ("sie kommt aus einer Library")
        return userError{name: userName}
    }
    // ...
}
```

## panic

`panic` ist eine andere Art von Fehler bzw. mit Fehlern umzugehen. `panic` ist eine Funktion in Go. Wenn sie aufgerufen wird, bricht das Programm ab und gibt den Stacktrace aus - mit folgender Ausnahme:

Im Allgemeinen sollte man `panic` nicht verwenden.

Wenn zumindest eine Funktion im Stacktrace in einem `defer` die Funktion `recover()` aufruft, bricht das Programm nicht ab.

Und nach Ansicht des Kursleiters sollte man `recover()` noch weniger verwenden.

Seine Empfehlung:

- "Normale" Fehler sollten mit dem error Return behandelt werden (die aufrufende Funktion muss den error wie "jeden anderen" Rückgabewert behandeln).
- Wenn etwas auftritt, was man nicht behandeln kann, sollte ein `log.Fatal` verwendet werden, um eine Meldung auszugeben und anschließend sollte das Programm beendet werden.

# Kapitel 7: Loops (for)

In Go gibt es nur eine `for` loop. Sie hat folgenden prinzipiellen Aufbau:

```golang
for INITIAL; CONDITION; AFTER {
  // eigentliche loop
}
```

INITIAL wird dabei initial einmal vor dem ersten Lauf der eigentlichen Loop ausgeführt. Dieser Schritt dient dazu gegebenenfalls Variablen zu initialisieren.
CONDITION wird **vor** jedem Schleifendurchlauf geprüft (aber **nach** dem Initialisieren), d. h. wenn die Bedingung am Anfang nicht erfüllt ist, wird die Schleife nie durchlaufen. Wenn in INITIAL Werte aus dem übergeordneten Block verändert werden, findet diese Veränderung auch statt, wenn die CONDITION nicht erfüllt ist.
AFTER wird **nach** jedem Schleifendurchlauf durchgeführt, das wird normalerweise verwendet um eine Zählvariable zu erhöhen oder zum nächsten Element in einer Liste zu gehen.

Es müssen aber nicht alle Teile angegeben werden - wenn beispielsweise CONDITION nicht angegeben ist hat man eine Endlosschleife produziert (sofern nicht im Schleifenkörper selbst die Schleife abgebrochen wird - z. b. mit `return` in einer `func`)

## while Schleife

Wie bereits festgestellt hat Go nur das Keyword `for` um eine Schleife zu schreiben. Eine Schleife, die sich wie eine sonst übliche while Schleife verhält, erreicht man dadurch, dass man der `for` Schleife nur eine CONDITION mitgibt. In diesem Fall darf man die `;` **nicht** schreiben (wenn man es doch tut, bekommt man einen Compilerfehler).

## "normalen" Schleifenlauf verändern

Man kann (in einer Funktion) nicht nur return dazu verwenden, um den Schleifenlauf zu verändern, sondern man kann auch `continue` und `break` verwenden. `continue` bricht den aktuellen, eigentlichen Loop ab und fährt mit der AFTER Anweisung fort. Das wird oft für **guard clauses** vewendet. `break` beendet die aktuelle Iteration und verlässt die gesamte Schleife, d. h. auch nachfolgende Iterationen werden nicht mehr ausgeführt. Das kann vor allem bei einer "**while**" Schleife praktisch sein.

# Kapitel 8: Slices

Arrays in Go sind eine fixe Anzahl von Elementen des gleichen Typs. Zur Anlage des Arrays muss die Anzahl der Elemente bekannt sein und kann sich nachträglich nicht mehr ändern. Bei Arrays wird die Anzahl der Elemente in [] dem Typ vorangestellt `[5]int` bezeichnet also ein Array, dass aus genau 5 `int` besteht.

Wenn man die Elemente dann auch gleich initialisieren will, geschieht dies wieder mit folgenden geschwungenen Klammern `{ }`. Die Elemente innerhalb dieser Klammern werden durch Beistriche `,` getrennt.

`primes := [6]int{2,3,5,7,11,13}`

Der größte Nachteil von `array`s in Go ist, dass sie eine fixe, unveränderliche Größe haben. Daher wird man fast immer stattdessen mit `slices` arbeiten.

Ein `slice` ist eine flexible Ansicht mit einer dynamischen Größe auf ein `array`. Man kann ein slice aus einem Array erstellen. Dabei gibt man den ersten und den letzten Index an, der enthalten sein soll. Dabei ist der letzte Index "gerade" **\*nicht** mehr enthalten\*!

```golang
primes := [6]int{2,3,5,7,11,13}
mySlice := primes[1:4] // ergibt [3, 5, 7]
```

Die generelle Syntax lautet `arrayname[ersterIndexEnthalten:letzterIndexNichtEnthalten]`. Sowohl ersterIndexEnthalten als auch letzerIndexNichtEnthalten kann weggelassen werden. Wenn der erste Index weggelassen wird, beginnt der slice beim ersten Element, wenn der zweite Index weggelassen wird, endet der slice nach dem letzten Element (in dem Fall ist also das letzte Element im Array noch im slice enthalten). Wenn beide Indizes weggelassen werden `[:]` umfasst der Slice also das gesamte Array.

Wenn man ein Slice einem anderen zuweist, verweisen beide auf das gleiche zugrundeliegende Array. Damit führt eine Änderung der Werte in einem Slice zur gleichen Änderung im anderen slice. In Wahrheit wird das zugrundeliegende Array verändert - und beide slices schauen auf das gleiche Array!

Weil slices nur eine "Ansicht" auf das zugrundeliegende Array sind, sind sie Referenzen. Daher ist es auch so, wenn man einer Funktion ein Slice als Parameter übergibt und diese Funktion den slice verändert, wird das zugrundeliegende Array verändert. Und die Änderung an dem Slice bzw. dem zugrundeliegenden Array ist nach der Funktion immer noch "sichtbar" bzw. existent!

Der **zero value** eines slice ist `nil`.

## make Funktion

Diese wird normalerweise für die Erzeugung eines slices verwendet - das "aus einem Array" schneiden ist die Ausnahme, weil Arrays schon die Ausnahme sind.

Die Funktion make hat dabei folgende Signatur: `func make([]T,len int,cap int) []T`, `len` ist dabei die initiale Länge des Slice und `cap` die Kapazität. Wobei die Kapazität nicht starr ist, sondern sich ändern kann. Wenn man an die Kapazitätsgrenze stößt erstellt Go automatisch ein neues Slice, dass eine doppelt so große Kapazität hat wie das aktuelle Slice. Daher kann man die Kapazität auch weglassen - in dem Fall wird sie gleich der Länge gesetzt, d. h. wenn nur ein Element dazu kommt muss Go das ganze Slice kopieren. Daher kann es trotzdem sinnvoll sein, eine Kapazität anzugeben (wenn man weiß/vermutet, wie groß das Slice "am Ende" sein wird).

Um beispielsweise ein Slice mit int zu erzeugen ist folgender Aufruf notwendig: `make([]int,5,10)`.

`make` füllt die Elemente mit dem **zero value** des Typs.

Um das Slice mit bestimmten Werten zu initialisieren, kann man ein slice literal verwenden. Nach der Initialisierung ist es aber dann ein ganz normales slices `mySlice:=[]string{"one","two","three"}`

## Länge und Kapazität eines slice

Die Länge eines Slice kann man mit der Funktion `len()` ermitteln und die Kapazität mit der Funktion `cap()`.

## variadic

Viele Funktionen können eine willkürliche Anzahl an Parametern für ihren letzten Parameter annehmen. Das wird vor allem von Funktionen in Standardbibliotheken oft verwendet. In der Funktionssignatur kennzeichnet man das mit `...` vor dem Typ, z. b. `func concat (strs ...string) string {}`. Diese Funktion nimmt eine beliebige Anzahl an strings entgegen und liefert einen einzigen string als Ergebnis zurück. In der Funktion selbst ist strs dann ein `slice` (von `string`s). Der Unterschied ist auf der Seite des Funktionsaufrufers - er übergibt eine variable Anzahl von `string`s (aber eben kein slice von `string`s).

## Spread Operator

Der Spread Operator (ebenfalls mit `...` geschrieben - er kommt aber **NACH** dem zu spreadenden Objekt - oft ein slice) erzeugt aus dem Slice eben für jedes Element im Slice einen eigenen Wert. Man kann ein int slice **nicht** an eine Funktion als Parameter übergeben, die sich ...int als Parameter erwartet, aber man kann slice... an diese Stelle übergeben (wenn der slice den Namen slice hat).

## append Funktion

Die append Funktion wird verwendet um dynamisch Elemente zu einem slice hinzuzufügen. `func append(slice []Type, elems ...Type)` []Type`. Die append Funktion hat also als ersten Parameter einen slice von einem bestimmten Typ und danach beliebig viele einzelne Werte mit diesem Typ. Und die Funktion liefert wiederum ein slice von diesem Typ zurück.

Wenn das zugrundeliegende Array nicht groß genug ist, wird ein neues, größeres Array erstellt und der returnierte slice zeigt auf dieses neue Array. Wenn das zugrundeliegende Array groß genug ist, werden die Elemente im bestehenden Array angefügt (in diesem Fall sieht man sie auch im als Parameter übergebenen slice!). Will man also sicherstellen, dass das übergebene Slice unverändert bleibt muss man es davor selbst kopieren! Das geht am einfachsten, indem man an ein leeres Slice das zu kopierende slice appended: `neu:=append([]int{},bestehendesSlice...)`.

Da die Funktion append einerseits das zugrundeliegende Array verändert und andererseits ein neues slice zurückgibt, sollte man append nur auf "sich selbst" anwenden, d. h. `someSlice = append(otherSlice, element)` sollte man vermeiden! (das "verbietet" damit den Trick aus dem vorigen Absatz - aber das Kopieren eines slice in ein anderes kann trotzdem interessant sein)

## range

`range` stellt eine Vereinfachung dar, wenn man in einem `for` über alle Elemente eines `slice` iterieren will:

```golang
for INDEX, ELEMENT := range SLICE {

}
```

ELEMENT ist dann das jeweilige Element aus dem slice SLICE. INDEX ist der Index dieses Elements (beginnt bei 0). Wenn man für den INDEX keine Verwendung hat, verwendet man einfach den blank operator `_` statt dessen.

## slice of slices

Slices können andere slices enthalten - wodurch man eine 2d slice erzeugen kann.

rows := [][]int{}

## Currying

Funktions currying ist ein Konzept aus der funktionalen Programmierung. Es erlaubt es, eine Funktion mit mehreren Parametern in eine Serie von Funktions mit jeweils nur einem Parameter "umzuwandeln".

Ein Beispiel in Go:

```golang
func main() {
  squareFunc := selfMath(multiply)
  doubleFunc := selfMath(add)

  fmt.Println(squareFunc(5))
  // prints 25

  fmt.Println(doubleFunc(5))
  // prints 10
}

func multiply(x, y int) int {
  return x * y
}

func add(x, y int) int {
  return x + y
}

func selfMath(mathFunc func(int, int) int) func (int) int {
  return func(x int) int {
    return mathFunc(x, x)
  }
}
```

Die Funktion selfMath nimmt eine Funktion mit zwei Parametern als Parameter entgegen und liefert eine Funktion mit einem Parameter zurück. Die zurückgelieferte Funktion ist die übergebene nur dass der Eingabeparameter (der zurückgelieferten Funktion) für beide Parameter der übergebenen Funktion verwendet wird. Darum wird - wenn man selfMath die add Funktion übergibt eine Verdopplungsfunktion (p + p) zurückgegeben. Wenn man selfMath die multiply Funktion übergibt, wird eine sqr Funktion (p \* p) zurückgegeben.

Allgemeiner gesprochen bezeichnet curryng eine Funktion, die eine Funktion als Parameter entgegennimmt und selbst wieder eine Funktion zurückliefert. Die currying Funktion erweitert quasi die übergebene Funktion.

In Go wird currying beispielsweise oft für Middleware eingesetzt - etwas, was bei jedem Request zusätzlich passieren soll.

# Kapitel 9: Maps

Maps sind vergleichbar mit Objekten in Javascript. Maps stellen eine Sammlung von key-value Paaren dar. In einem Map kann jeder key nur ein einziges Mal vorkommen.

Der **zero value** von Maps ist `nil`.

Maps sind - wie slices - Referenzen auf "darunterliegende" Daten, d. h wenn einer Funktion eine Map übergeben wird und diese ändert die Map, ist diese Änderung auch nach dem Ende der Funktion "gültig" bzw. "vorhanden".

## Erzeugen von Maps

Maps können mit Hilfe der `make` Funktion erstellt werden. Dies sieht ähnlich wie für ein slice aus: `make(map[string]int)` Erstellt ein Map, wo der key Datentyp `string` ist und der value Datentyp `int` ist.

Man kann auch ein Map Literal erstellen: `map[string]int{"John":36,"Mary":21}` (erstellt ein Map, wo der key ein `string` ist und der value ein `int`. Und es werden auch gleich zwei Elemente "John" mit 36 und "Mary" mit 21 angelegt).

## Mutationen von einem Element eines Maps

### Einfügen / Ändern

`m[key] = elem`

### Auslesen eines Elements

`elem = m[key]`

Achtung: Hier bekommt man immer etwas zurück - entweder den Wert des Elements mit dem Key key oder den zero value für den Datentyp von Map.

### Element löschen

`delete(m,key)`

### Prüfen, ob ein key in einem Map vorhanden ist

`elem, ok := m[key]`

Wenn ein Element mit dem Key key vorhanden ist, ist elem dieses Element und ok ist true. Wenn es nicht vorhanden ist, ist elem der zero value der Elemente im Map und ok ist false. Dieses ", ok" hat den Namen **comma ok**.

### Was ändert sich, wenn elem kein Basistyp ist (z. b. ein struct)

Dann scheinen die Dinge unschön zu werden, man kann nämlich ein Attribut des structs nicht einfach mit dot-Notation setzen, sondern man muss einen Umweg gehen:

```golang
type user struct {
  name: string
  age: int
}

func main() {
  users := make([string]user)
  users["John"] = {name:"John",age:27}
  users["Bob"] = {name:"Bob",age:23}

  // COMPILERFEHLER !
  // users["Bob"].age = 43

  // korrekt
  b := users["Bob"]
  b.age := 43 // nach dieser Anweisung ist das Alter von users["Bob"] immer noch 23
  users["Bob"] = b
}
```

Mit der dotNotation kann man die Werte aus dem Struct nur auslesen (`fmt.Printf("%v Jahre alt\n",users["Bob"].age)` funktioniert und liefert wie erwartet "43 Jahre alt")

## Was darf ein Key in einer map sein?

Als value kann alles in einer map vewendet werden. Für die Keys trifft das nicht zu: **map keys may be of any type that is comparable**. Was sind jetzt Typen die compareable sind? Alles, was mit `==` verglichen werden kann, das sind `boolean`, `numeric`, `string`, `pointer`, `channel`, `interface types` und structs oder arrays, die nur die vorher genannten Typen enthalten. Insbesondere bemerkenswert ist, dass slices, maps und functions **nicht** auf der Liste stehen - und damit als keys für maps nicht verwendet werden können.

Interessant ist, dass ein `struct` als key verwendet werden kann - damit kann man leicht mehrdimensionale Maps erzeugen. Nehmen wir an, wir wollen tracken, wie oft jede Seite auf unserem Server von jedem Land aufgerufen wurde. Das kann man relativ leicht mit einem entsprechenden Map mit einem struct als Key machen:

```golang
type Key struct {
  Path, Country string
}

hits := make(map[Key]int)

hits[Key{"/doc","at"}]++ // wenn jemand aus Österreich die /doc Seite besucht hat - wenn der Eintrag in hits noch nicht vorhanden ist, wird er angelegt
hits[Key{"/help","de"}]++ // wenn jemand aus Deutschlang die /help Seite besucht hat

fmt.Printf("%v Zugriffe aus Deutschland auf /help",hits[Key{"/help","de"}]) // Gibt die Anzahl der Zugriffe auf die /help Seite aus Deutschland aus
```

Anmerkung: Mit verschachtelten maps wäre das wesentlich aufwändiger (weil beispielsweise das automatische Anlegen wenn eine neue Seite angesurft wird nicht passieren würde - weil der Aufruf ja wäre diese Seite wurde gerade aus einem bestimmten Land angesurft und dann müsste man sich selbst darum kümmern, ob für die Seite schon eine "äußere" Map vorhanden ist und wenn nicht diese auch selbst anlegen).

# Kapitel 10: Pointer

Ein Pointer ist eine Variable, die die Adresse einer anderen Variable speichert. Ein Pointer wird mit `*` erstellt. `var p *int` - p ist ein Zeiger auf einen int.

Der `&` Operator wiederum liefert die Adresse der Variablen zurück, nicht den Wert selbst.

```golang
x := 5; // x wird beispielsweise an Adresse 101 gespeichert - in Adresse 101 steht dann 5
p := &x; // p wird beispielsweise an Adresse 105 gespeichert - und darin steht dann 101
```

Und wenn man einen Pointer an eine Funktion übergibt (statt dem Wert, der an der Stelle steht, wo der Pointer hinzeigt) und die Funktion den Wert, auf die der Pointer zeigt, ändert - dann ist diese Änderung auch nach dem Ende der Funktion gültig (weil die Funktion keine Kopie des Wertes erhalten hat, sondern einen Verweis auf die eigentliche Adresse).

Um dann (über p) auf den Wert von x zuzugreifen, verwendet man ebenfalls den `*` Operator (**Dereferenzierungsoperator**), d. h. *p liefert im vorigen Beispiel 5. Und wenn man *p = 27 ausführt, wird an die Stelle im Speicher, wo z hinzeigt, 27 geschrieben. Und weil die Variable x diese Stelle im Speicher verwendet, hat sie jetzt auch den Wert 27.

```golang
func info(i *int) {
  fmt.Printf("Info %v\n",*i)
}

z := 17
info(&z) // ergibt Info 17
```

## Zeiger auf struct

```golang
type coord struct {
  x int
  y int
}

func printCoords(p *coord) {
  fmt.Printf("(%v/%v)\n",(*p).x,(*p).y) // die Klammer rund um *p ist NOTWENDIG, sonst bekommt man einen Compilerfehler, weil scheinbar . einen höheren "Rang" wie * hat
}

c := coord{x:12,y:-6}
printCoords(&c) // ergibt (12/-6)
```

Für Zeiger auf struct gibt es eine "Abkürzung", man muss den * nicht angeben - und kann dann auch die `()` weglassen, d. h. die Zeile in printCoords kann man auch so schreiben: `fmt.Printf("(%v/%v)\n",p.x,p.y)`. Go erkennt, dass p ein Zeiger ist und ergänzt (*p) automatisch - ohne es in den Programmcode zu schreiben.

Die Klammer um _p ist notwendig, weil der . Operator eine höhere Priorität wie der _ Operator hat, d. h. _p.x wird als _(p.x) interpretiert - und das ist nicht zulässig, weil .x kein Pointer ist, sondern ein Wert.

## Leerer Pointer

Man kann einen "leeren" Pointer definieren, also einen Pointer der auf keine Adresse zeigt. In diesem Fall hat der Pointer den Wert `nil`. Diese Zeiger nennt man auch **nil pointer**. Diese sind sehr "gefährlich", wenn man versucht einen nil pointer zu dereferenzieren bricht das Programm mit einer panic ab! Daher sollte man _immer_ bevor man einen Pointer dereferenzieren will prüfen, ob er nicht nil ist.

## Pass by reference

In Go werden nicht zusammengesetzte Typen als **call by value** übergeben, d. h. die Funktion erhält eine Kopie des Parameters und d. h. dass Änderungen an den Werten solcher Parameter innerhalb der Funktion in der Außenwelt "unsichtbar" sind.

Einer der häufigsten Anwendungsfälle von Pointern in Go ist, dass man Parameter **by reference** übergeben will. Damit erhält die Funktion keine Kopie sondern die Adresse der Daten - und wenn sie dann an dieser Adresse Änderungen am Wert vornimmt, bleiben die auch nach ihrer Laufzeit erhalten (und sind somit für die Außenwelt sichtbar).

## Pointer receiver

Ein receiver Typ bei einer Methode kann ein Pointer sein. In dem Fall kann die Methode die Daten im struct verändern. Da dies ein häufiger Anwendungsfall für Methoden ist, sind pointer receiver üblicher als normale receiver.

Wenn man eine Methode aufruft, muss man sich auch keine Gedanken darüber machen, ob sie mit einem normalen oder einem Pointer receiver definiert ist. Man kann immer die Variable direkt verwenden. Der Pointer wird automatisch abgeleitet.

```golang
type car struct {
  color string
}

func (c *car) setColor(color string) {
  c.color = color
}

func (c car) dontSetColor(color string) {
  c.color = color
  fmt.Printf("color inside of dontSetColor=%v\n",c.color)
}

func main() {
  c := car{
    color: "white",
  }
  c.setColor("blue")
  fmt.Println(c.color) // "blue"
  c.dontSetColor("red") // "color inside of dontSetColor=red"
  fmt.Println(c.color) // "blue
}
```

Die Methode setColor ändert die Farbe (weil sie einen Pointer Receiver verwendet), die Methode dontSetColor ändert die Farbe nicht dauerhaft (außerhalb der Methode), weil sie einen Value Receiver verwendet. Im Aufruf unterscheiden sich die beiden Methoden aber nicht. Man könnte auch `(&c).setColor("blue")` verwenden - das Ergebnis wäre das gleiche, aber ohne ist es einfacher zu Schreiben und auch zu lesen (daher wird man de facto nie Code finden, der diese längere Form verwendet).

## Pointer Performance

Es ist nicht notwendig/sinnvoll überall Pointer zu verwenden um die Zeit für das Kopieren der Parameter beim Funktionsaufruf einzusparen. Es ist wichtiger klaren, korrekten und wartbaren Code zu schreiben. Wenn man ein Performanceproblem hat, muss man dieses lösen (und ja, da kann es dann sinnvoll/notwendig sein an den richtigen Stellen Pointer zu verwenden - aber andere Dinge sollten zuerst bedacht werden).

Stack vs. Heap - die lokalen Variablen werden im Stack gespeichert, der schneller im Zugriff ist als der Heap (wo die Pointer ins Spiel kommen). Wenn der Wert so groß ist, dass das Kopieren einen Performanceeinfluss hat, kann es tatsächlich zielführend sein in diesem Fall einen Pointer zu verwenden.

# Kapitel 11: Packages

Jedes Go Programm besteht aus Packages. Es hat ein main package. In diesem gibt es eine `func main()`. Diese ist der Einsprungspunkt in das Programm - diese Funktion wird ausgeführt, wenn das Programm gestartet wird. Alle anderen Dateien sind library packages und dürfen **KEIN** main() haben. Aus diesen anderen Dateien werden nur Dinge exportiert - aber nichts wird direkt ausgeführt.

Library packages haben per Konvention den letzen Teil ihres Pfades als Package name.

In einem Verzeichnis darf es nur ein Package geben. Wenn unterschiedliche Dateien in einem Verzeichnis zu unterschiedlichen Packages gehören führt dies zu einem Fehler. Und das gilt sowohl für main als auch library packages.

## Module

Go Programme sind in Packages organisiert. Ein Package ist dabei ein einzelnes Verzeichnis. Alles was in einem Go File in einem Verzeichnis vorhanden ist, ist für alle anderen Go Files im gleichen Verzeichnis sichtbar.

Ein Repository kann ein oder mehrere Module enthalten. Ein Modul ist eine Sammlung von Go Programmen, die gemeinsam released werden. Trotzdem ist es üblich ein Modul pro Repository zu haben.

Ein Modul besteht aus einer go.mod im Rootverzeichnis des Projekts. Es enthält den `module` Pfad, die `go` version und es kann Verweise auf andere benötigte externe Pakete haben.

```golang
module github.com/powiedl/gocourse/exampleproject

go 1.24.0

require github.com/google/examplepackage v1.3.0
```

Der Modulpfad gibt dabei auch an, wo man das Modul herunterladen kann (im obigen Beispiel würde go das Modul unter https://github.com/powiedl/gocourse/exampleproject erwarten).

## Go Umgebung

- Im Normalfall hat man viele git repositories auf seinem System (üblicherweise eines pro Projekt)
- Jeses Repository ist normalerweise ein einzelnes Modul
- Jedes Repository enthält ein oder mehrere Packages
- Jedes Package besteht aus einem oder mehreren Go source files in einem einzelnen Verzeichnis. Und in einem Verzeichnis darf es nur ein Package geben.

## GOPATH

Früher hat man die Umgebungsvariable GOPATH verwendet - und darin hat man den Sourcecode in einem Verzeichnis `src` gehabt. Das ist mittlerweile nicht mehr notwendig bzw. üblich.

## Programm initialisieren

Man wechselt in das Verzeichnis, in dem man sein Programm schreiben will. Dort gibt man auf der Kommandozeile den folgenden Befehl ein `go mod init {REMOTE}/{USERNAME}/programmname` {REMOTE} ist dabei der Remote Source Provider (z. b. `github.com`), {USERNAME} der Benutzername dort (z. b. `powiedl`) und programmname eben der Programmname. Sinnvollerweise ist er gleich wie der Verzeichnisname, in dem man sich gerade befindet.

Der Befehl legt an dieser Stelle ein go.mod an (das so aussieht, wie im Punkt [Module](#module) beschrieben.

Außerdem legt man in diesem Verzeichnis ein `main.go` mit folgendem Aufbau an:

```golang
package main

func main() {

}
```

Dieses Programm (und natürlich hat man sinnvollerweise auch einen Inhalt in `main()`) kann man dann mit `go run main.go` ausführen. Dabei wird aber kein Executeable gebaut, sondern es wird nur im Memory das ausführbare Programm erstellt und aufgerufen. Dieses Kommando eignet sich daher in erster Linie nur zum schnellen Testen von kleineren Skripten/Programmen.

## Go build

Mit `go build` wird ein ausführbares Programm aus der angegebenen Datei erzeugt. In diese Datei wird alles eingefügt, was man benötigt um das Programm ausführen zu können, das bedeutet, dass es reicht das erzeugte File auf einem anderen Computer (mit "gleichem" Betriebssystem) auszuführen. Man muss sonst nichts zur Verfügung stellen und auf dem anderen Computer muss auch nichts extra installiert werden (oder vorhanden sein).

## Go install

Mit `go install` in dem Repoverzeichnis wird das Programm kompiliert und das Executeable wird im GOBIN directory erstellt (und das GOBIN Directory ist im Normalfall im Pfad enthalten, damit kann man es dann von überall aus ausführen).

## Was ein Package exportiert?

Alles was mit einem Großbuchstaben beginnt, wird von einem Package exportiert, alles andere ist nur privat innerhalb des Packages verfügbar. Man muss also nicht extra irgendo exports oder ähnliches angeben - das wird einzig und allein nach dem ersten Zeichen der Objektname bestimmt.

Um etwas aus einem Package in seinem Programm zu verwenden, muss man einerseits das Package importieren und andererseits der zu verwenden Funktion den Packagenamen voranstellen. Am einfachsten ist es an einem Beispiel zu erkennen:

```golang
package main
import (""fmt";"github.com/powiedl/mystrings")

func main() {
  fmt.Println(mystrings.Reverse("hallo Welt"))
}
```

Im Beispiel wird also die Funktion Reverse (man beachte das große R am Anfang des Namens) auf den string "hallo welt" angewandt und das Ergebnis dieser Funktion mittels fmt.Println ausgegeben.

## Lokales übersteuern von importierten requires

In der `go.mod` des jeweiligen Packages ist auch angegeben, welche Dependencies dieses Package hat. Im Beispiel unter [Module](#module) ist es `github.com/google/examplepackage v1.3.0`. Während der Entwicklung und der Verwendung von eigenen Packages kann es zielführend sein, wenn man den offiziellen Pfad "überschreibt". Dies macht man, indem man vor dem require eine (oder mehrere) replace Anweisungen in `go.mod` einfügt.

```golang
replace github.com/google/examplepackage v1.3.0 => ../examplepackage

require github.com/google/examplepackage v1.3.0
```

Damit sagt man dem Go Compiler, dass er statt dem offiziellen examplepackage auf github das Package aus dem "Nachbarverzeichnis" examplepackage verwenden soll.

Das funktioniert allerdings nur in einer lokalen Entwicklungsumgebung.

## Go get

Mit `go get {REPOSITORYLOCATION}/{USERNAME}/{REPOSITORYNAME}` lädt und installiert man ein externes Package. Dieses wird im lokalen Package cache gespeichert und steht ab sofort allen Go Programmen zur Verfügung. Im Verzeichnis, wo das eigene Programm liegt, wird eine `go.sum` Datei erstellt. In dieser trackt Go sämtliche Abhängigkeiten von dem Programm (und die Abhängigkeiten der Abhängigkeiten usw).

## Regeln für "saubere" Packages

1. **Hide internal logic**: Man sollte nur die Teile exportieren (mit einem großen Anfangsbuchstaben beginnen lassen), die für die Interaktion mit dem Package absolut notwendig sind. Alles, was das Package darüber hinaus für das eigene Funktionieren benötigt, sollte nicht exportiert werden.
2. **Don't change APIs**: Die exportierten Teile sollten (von den Signaturen) so stabil wie irgendwie möglich gehalten werden und sollten sich nur in Ausnahmefällen ändern. Für die Verwender des Packages gibt es nichts unangenehmeres, wie wenn sie jedes Mal, wenn sie eine neue Version des Packages verwenden, etwas an ihrer Verwendung des Packages ändern müssen.
3. **Don't export functions from the main package**: Man soll keine Funktionen aus dem main Package exportieren, weil es keine Library ist.
4. **Packages shouldn't know about dependents**: So wie ein Interface nichts über die structs wissen soll, die es implementieren soll ein Package nichts über seine Verwender wissen (müssen).

# Kapitel 11 Concurrency

Concurrency bezeichnet die Fähigkeit mehrere Codeteile "gleichzeitig" auszuführen. Bei einem Single Core System schaltet das Betriebssystem sehr schnell zwischen den einzelnen Teilen hin und her wodurch der Eindruck entsteht, dass die einzelnen Teile gleichzeitig passieren. Bei einem Multicoresystem können wirklich mehrere Dinge gleichzeitig ausgeführt werden (eben ein Teil pro Core) - in diesem Fall spricht man von Parallelität. Im Sprachgebrauch wird beides "vermischt" eingesetzt.

Go wurde von Anfang an dafür designed Teilaufgaben parallel auszuführen. Dazu reicht es vor dem Aufruf der Funktion, die man parallel ausführen will `go` voranzustellen. Das `go` Keyword startet eine neue **Go Routine**. Die aufrufende Funktion wartet nicht auf das Ende dieser Go Routine sondern macht sofort mit der nächsten Anweisung weiter! Man kann daher nicht auf das Ergebnis der Go Routine warten indem man den Rückgabewert der Go Routine entgegennimmt. In Wahrheit macht es keinen Sinn, dass eine Go Routine einen Rückgabewert hat (es kann nur sein, dass man eine bestehende Funktion - die einen Rückgabewert hat - "parallel" ausführen will - aber dann kann man eben nicht auf den Rückgabewert dieser Funktion zugreifen).

## Channels

Channels sind eine typisierte, thread-sichere Queue. Über diese Channels können verschiedene Go Routinen miteinander kommunizieren. Über diese Channels kann man dann auch die "Rückgabewerte" von Funktionen zwischen unterschiedlichen Go Routinen austauschen. Die Channels sind sowohl lesend als auch schreibend blockierend, d. h. wenn eine Funktion einen Wert aus einem leeren Channel lesen will wartet sie solange, bis ein Wert im Channel auftaucht. Und wenn eine schreibende Funktion in einen vollen Channel schreiben will, wartet sie solange, bis der Channel nicht mehr voll ist (weil ein Reader eine Information entnommen hat).

### Erzeugen eines Channels

Wie maps und slices müssen auch Channels erzeugt werden, bevor sie benutzt werden können. Dabei wird auch der Typ festgelegt, den der Channel aufnehmen kann. Dazu dient die Funktion `make` und das Keyword `chan`: `ch := make(chan int) // erzeugt einen Channel vom Typ int mit dem Namen ch`

Ein deklarierter, aber nicht initialisierter Channel hat den Wert nil (`var c chan string // c == nil`).

Sowohl das Lesen aus einem nil channel als auch das Schreiben in einen nil channel blockiert für immer.

### Daten an einen Channel senden

Im wesentlichen gibt es zwei Operationen für einen Channel. Man kann Daten (vom richtigen Typ) hineinsenden oder man kann diese Daten daraus empfangen.

`ch <- 69 // Sendet 69 in den zuvor erzeugten Channel`

Der `<-` wird als Channel Operator bezeichnet und er soll die Datenflussrichtung darstellen (in dem Fall wird etwas in den Channel gesendet). Werden mehrere Werte in den Channel geschrieben, so behalten sie dort ihre Reihenfolge, d. h. der als erstes geschriebene Wert steht ganz am Anfang des Channels und danach der als zweite geschriebene usw. Das ist beim Empfangen wichtig.

### Daten aus einem Channel empfangen

`v := <- ch // Die Variable v wird initalisiert mit dem Wert, der im Channel "wartet"`

Danach wird der Wert aus dem Channel entfernt. Wenn kein Wert im Channel vorhanden ist, wartet das Programm an dieser Stelle so lange, bis ein Wert in den Channel geschrieben wird.

### Empty Channels

Manchmal ist es egal, was durch einen Channel geschickt wird, es geht nur darum, dass etwas geschickt wurde. Dieses etwas kann man dann als "ich bin fertig" vom Sender interpretieren. Um das auszudrücken, verwendet man das Empfangen aus einem Channel ohne "Empfänger": `<- ch`. Durch diese Anweisung wird die Funktion an dieser Stelle angehalten, solange bis etwas aus dem Channel ch empfangen wird. Das was empfangen wird, wird sofort verworfen (es geht also nur darum, dass etwas empfangen wurde).

## Buffered Channels

Channels können optional gebuffert werden, d. h. sie haben Platz für mehr als ein Element. Bei einem ungebufferten Channel kann nichts mehr in den Channel gesendet werden, wenn sich ein Element darin befindet. Bei einem gebufferten Channel können sich bis zur Buffergröße Elemente im Channel befinden. Erst dann ist es nicht mehr möglich etwas in den Channel zu senden. Wenn eine Go Routine etwas in einen Buffer senden will, aber darin kein Platz ist, wartet sie solange, bis irgendetwas ein Element aus dem Channel entnimmt (womit wieder Platz für ein neues Element wird).

Die Größe wird der `make` Funktion als optionaler zweiter Parameter übergeben.

## Schließen eines Channels

Der Sender kann den Channel explizit schließen. `close(ch) // Damit wird der Channel ch geschlossen`. Danach kann man nichts mehr in den Channel schicken. Wenn man es doch macht, kommt es zu einer **panic**. Wenn man aus einem geschlossenen Channel liest, erhält man sofort den **zero value** für den Datentyp des Channels als Ergebnis.

### Kontrolle, ob ein Channel geschlossen ist

Die Empfänger können prüfen, ob ein Channel geschlossen ist. Dazu gibt es beim Lesen des Channels einen zweiten Return value vom Typ `bool`. Wenn der Channel geschlossen ist, hat dieser den Wert `false`. Ein Channel wird vom Sender geschlossen, solange sich aber noch Werte im Channel befinden, die kein Empfänger empfangen hat, ist der Channel für das Auslesen noch offen. Darum ist es "sicher", wenn man beim Auslesen prüft, ob der Channel geschlossen ist.

## range und Channels

Man kann über Channels rangen - so wie bei Slices und Maps. Wenn nichts im Channel ist wartet range auf das nächste Element und blockiert die Verarbeitung solange. Wenn der Channel geschlossen wird, bricht range das for ab und das Programm fährt mit der Anweisung nach dem `for` Block fort. `for item := range ch { ... }`

Das erleichtert das Arbeiten mit Channels "ungemein".

## select

Ein `select` Statement wird verwendet um gleichzeitig auf mehrere Channels zu hören. Es ähnelt einem `switch` Statement (ist aber für verschiedene Channels gedacht).

```golang
select {
  case i,ok := <- chInts:
    fmt.Println(i)
  case s,ok := <- chStrings:
    fmt.Println(s)
}
```

Für den Channel, der einen Wert in sich ready hat, wird der case Block entsprechend ausgeführt. Wenn mehrere Channel gleichzeitig einen Wert ready haben wird zufällig aus diesen Channels **einer** gewählt und dessen Block wird ausgeführt. Im Normalfall befindet sich das select in irgendeiner Schleife, d. h. beim nächsten Durchlauf wird dann ein anderer case Block ausgeführt. Wenn das `select` aber wirklich nur einmal ausgeführt wird, bleiben die Werte in den anderen Channels "ewig" stehen.

### default case

Man kann in einem select auch einen default case angeben. Dieser wird ausgelöst, wenn keiner der überwachten Channels einen Wert ready hat. Er löst damit die Blockade, die das `select` sonst auf die umgebende Funktion hat.

## Tickers

Ticker sind spezielle Channels aus der `time` Standardbibliothek. Diese drei Ticker Funktionen sind gebräuchlich:

- `time.Tick()` - liefert einen Channel, der einen Wert im angegebenen Intervall sendet (also wiederkehrend)
- `time.After()` - liefert einen Channel, der einmalig nach dem angegebenen Intervall einen Wert sendet und sich danach schließt
- `time.Sleep()` - Blockiert die aktuelle Go Routine für die angegebene Zeitdauer

Alle drei Funktionen nehmen eine time.Duration entgegen, das ist ein Wert mit einer Zeiteinheit (Z. b. `time.Millisecond`) multipliziert. Wenn keine Zeiteinheit angegeben wird handelt es sich um Nanosekunden (wobei für mich fraglich ist, ob man so wirklich 17 Nanosekunden warten kann ...).

## Readonly /Writeonly Channels

Wenn man Channels als Parameter übergibt, kann man festlegen, ob dieser Channel nur gelesen oder nur geschrieben werden darf - oder im Standardfall beides.

```golang
func readch(ch <-chan int) { ... } // hierbei handelt es sich um einen Read-Only channel für die Funktion readch
func writech(ch chan<- int) { ... } // hierbei handelt es sich um einen Write-Only channel für die Funktion writech
```

Diese explizite Klarstellung kann einerseits dem Leser helfen, zu verstehen was die Funktion mit dem Channel machen will und außerdem prüft der Go Compiler, ob man sich in der Funktion daran hält (d. h. ob man eh nicht versucht in einen Read-Only Channel zu schreiben oder umgekehrt). Es ist daher vorteilhaft immer so explizit wie möglich zu sein.

# Kapitel 13 Mutexes

Mutexe erlauben es den Zugriff auf Daten zu sperren. Damit können wir sicherstellen, welche Go Routine zu welcher Zeit auf bestimmte Daten zugreifen kann. In den Standardbibliotheken gibt es eine integrierte Implementierung von Mutexes - `sync.Mutex`.

Dazu stellt das zwei Methoden zur Verfügung: `.Lock()` und `.Unlock()`. Zwischen `.Lock()` und `.Unlock()` kann nichts anderes auf die geschützten Daten zugreifen. Es ist gute Praxis, diesen Block in eine eigene Funktion "auszulagern" - am Anfang dieser Funktion `.Lock()` und auch gleich `defer `.Unlock()` aufzurufen und die geschützten Anweisungen eben in dieser Funktion zu kapseln.

Das prinzipielle Problem, dass Mutexe lösen ist, wenn verschiedene Threads gleichzeitig auf eine geteilte Resource zugreifen wollen - wobei mindestens einer schreibend zugreifen will. Da kann es nämlich sein, dass ein lesender Thread veraltete Daten liest. Und in dem Fall wird das Go Programm eine **panic** auslösen.

## Maps sind nicht thread-safe

Wenn man mehrere Go Routinen hat, die auf die gleiche Map zugreifen und zumindest eine davon ändert auch die Map muss man die Map mit einem Mutex locken.

## Ein Beispiel

```golang
package main

import (
	"fmt"
)

func main() {
	m := map[int]int{}
	go writeLoop(m)
	go readLoop(m)

	// stop program from exiting, must be killed
	block := make(chan struct{})
	<-block
}

func writeLoop(m map[int]int) {
	for {
		for i := 0; i < 100; i++ {
			m[i] = i
		}
	}
}

func readLoop(m map[int]int) {
	for {
		for k, v := range m {
			fmt.Println(k, "-", v)
		}
	}
}
```

Dieses Programm erstellt eine Map und zwei Go Routinen. Eine schreibt permanent in die Map während die zweite gleichzeitig aus der selben Map liest. Auf einer Multicore-Maschine führt das zu einem "fatal error: concurrent map iteration and map write".

Um das zu lösen, muss man einen entsprechenden Mutex hinzufügen.

```golang
package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[int]int{}

	mu := &sync.Mutex{}

	go writeLoop(m, mu)
	go readLoop(m, mu)

	// stop program from exiting, must be killed
	block := make(chan struct{})
	<-block
}

func writeLoop(m map[int]int, mu *sync.Mutex) {
	for {
		for i := 0; i < 100; i++ {
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}
	}
}

func readLoop(m map[int]int, mu *sync.Mutex) {
	for {
		mu.Lock()
		for k, v := range m {
			fmt.Println(k, "-", v)
		}
		mu.Unlock()
	}
}
```

In diesem Beispiel wird der Schreibende und Lesende Zugriff auf die Map jeweils von einem Mutex "geklammert", dadurch kann nur entweder das Lesen oder das Schreiben stattfinden - und nicht beides gleichzeitig. Wichtig: Der Mutex muss dabei ein Pointer sein (darum wird er auch mit der Adresse von `sync.Mutex` initialisiert).

## RW Mutex

Die Standardlibrary hat auch einen RW Mutex, der zwei weitere Methoden (`.RLock()` und `.RUnlock()`) hat. Diese haben Performancevorteile bei Leseintensiven Operationen. Beliebig viele RLocks können gleichzeitig vorliegen - ohne sich gegenseitig zu stören. Aber es kann nur ein Lock aktiv sein (der andere Locks und RLocks blockiert).

# Kapitel 14: Generics

Generics sind "relativ" neu in Go. Obwohl Go keine Klassen kennt, hat mittlerweile das Konzept von Generics Einzug gehalten. Damit kann man den Typ von etwas wegabstrahieren - wenn es für die konkrete Aufgabe, die der jeweilige Code löst, keine Bedeutung hat. Ein Beispiel dafür wäre eine Funktion die ein slice in der Mitte in zwei halb so große slices teilt. Für diese Aufgabe ist es komplett egal, welcher Datentyp in dem slice gespeichert wird. Früher musste man denselben Code für alle benötigten Typen "1:1" wiederholen (eben nur den Typ austauschen).

Generics erlauben es für Typen Variablen zu verwenden. Wie auch bei Typescript gibt man das Generic nach dem Funktionsnamen an:

```golang
func splitAnySlice[T any](s []T) ([]T, []T) {
  mid := len(s)/2
  return s[:mid],s[mid:]
}
```

In diesem Beispiel hat man die Typvariable T definiert. Nach dem Funktionsnamen hat man angegeben, dass T keinen Einschränkungen unterliegt (weil es für die Funktion wirklich total egal ist, welchen Typ die einzelnen Elemente in dem Slice haben).

## Warum verwendet man Generics?

Generics helfen sich wiederholenden Code zu vermeiden oder zumindest zu reduzieren. Wenn sich eine bestimmte Funktionalität nicht unterscheidet, aber für unterschiedliche Datentypen benötigt wird, sind Generics die richtige Wahl. Man kann damit die Funktionalität nur einmal schreiben und für verschiedene Datentypen verwenden. Besonders in Libraries kann das von großem Vorteil sein.

## Constraints

Mit Constraints kann man gewisse Bedingungen definieren, die ein Typ erfüllen muss, damit er unter ein bestimmtes Generic "passt". Das kann manchmal notwendig/sinnvoll sein. Constraints sind einfach interfaces, die es erlauben, dass ein Generic eben nur für bestimmte Typen passt.

### Erstellen eines Constraints

Am offensichtlichsten ist es anhand eines Beispiels:

```golang
type stringer interface {
  String() string
}

fonc concat[T stringer](vals []T)string {
  result := ""
  for _,val := range vals {
    result += val.String() // hier wird die String Methode auf dem Generic verwendet, daher muss der Generic dieses Interface implementieren
  }
}
```

Im obigen Beispiel wird ein `interface` stringer definiert. Dieses besteht aus der Methode String(), die einen `string` zurückliefert. Die Idee dieser Methode ist, dass etwas, dass das stringer Interface implementiert eine Methode String hat, die dazu führt, dass sich das etwas als `string` darstellt.

Diese Methode String wird dann in der Funktion selbst aufgerufen, d. h. das interface, muss alle spezifischen Methoden enthalten, die in der Funktion für den Generic Datentyp benötigt werden.

## Interface Type Lists

Seit der Einführung von Generics gibt es auch Interface Type Lists. Damit kann man eine Liste von Typen zu einem "Übertypen" zusammenfassen:

```golang
type Ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
        ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
        ~float32 | ~float64 |
        ~string
}
```

Im Beispiel die Ordered Typen. Diese unterstützen die Vergleichsoperatoren: `<`, `<=`, `>` und `>=`

## Parametrisierte Constraints

Die Interface Definitionen, die als constraints verwendet werden, können selbst Typparameter haben. Im folgenden ein Beispiel für ein store interface, das ein Geschäft darstellt, dass verschiedene Produkte verkauft. Damit ein Produkt ein "gültiges" Produkt wird, muss es die Methoden Name() und Price() implementieren:

```golang
type store[P product] interface {
	Sell(P)
}

type product interface {
	Price() float64
	Name() string
}

type book struct {
	title  string
	author string
	price  float64
}

func (b book) Price() float64 {
	return b.price
}

func (b book) Name() string {
	return fmt.Sprintf("%s by %s", b.title, b.author)
}

type toy struct {
	name  string
	price float64
}

func (t toy) Price() float64 {
	return t.price
}

func (t toy) Name() string {
	return t.name
}

// The bookStore struct represents a store that sells books.
type bookStore struct {
	booksSold []book
}

// Sell adds a book to the bookStore's inventory.
func (bs *bookStore) Sell(b book) {
	bs.booksSold = append(bs.booksSold, b)
}

// The toyStore struct represents a store that sells toys.
type toyStore struct {
	toysSold []toy
}

// Sell adds a toy to the toyStore's inventory.
func (ts *toyStore) Sell(t toy) {
	ts.toysSold = append(ts.toysSold, t)
}

// sellProducts takes a store and a slice of products and sells
// each product one by one.
func sellProducts[P product](s store[P], products []P) {
	for _, p := range products {
		s.Sell(p)
	}
}

func main() {
	bs := bookStore{
		booksSold: []book{},
	}

    // By passing in "book" as a type parameter, we can use the sellProducts function to sell books in a bookStore
	sellProducts[book](&bs, []book{
		{
			title:  "The Hobbit",
			author: "J.R.R. Tolkien",
			price:  10.0,
		},
		{
			title:  "The Lord of the Rings",
			author: "J.R.R. Tolkien",
			price:  20.0,
		},
	})
	fmt.Println(bs.booksSold)

    // We can then do the same for toys
	ts := toyStore{
		toysSold: []toy{},
	}
	sellProducts[toy](&ts, []toy{
		{
			name:  "Lego",
			price: 10.0,
		},
		{
			name:  "Barbie",
			price: 20.0,
		},
	})
	fmt.Println(ts.toysSold)
}
```

# Kapitel 15: Enums

In Go gibt es keine Enums.

Das, was einem `Union Type` in Typescript am nächsten kommt ist:

```golang
type sendingChannel string // hier definieren wir einen Type Alias für string sendingChannel

const (
  Email sendingChannel = "email"
  SMS   sendingChannel = "sms"
  Phone sendingChannel = "phone"
)

func sendNotification(ch sendingChannel, message string) {
  // send the message
}
```

Damit verhindert Go, dass wir dieses Programm schreiben:

```golang
sendingCh := "slack"
sendNotification(sendingCh, "hello") // string is not sendingChannel
```

Aber es verhindert nicht, dass wir das schreiben: `sendNotification("slack","hello")` und es verhindert ebenfalls nicht folgendes:

```golang
sendingCh := "slack"
convertedSendingCh := sendingChannel(sendingCh)
sendNotification(convertedSendingCh, "hello")
```

## Iota

Iota ist ein Go Feature, dass in einem const block jeder Konstanten den nächsten int Wert zuweist. Es beginnt mit 0 bei der ersten Konstanten. Dort - und nur dort - wird auch das Keyword `iota` als Wert für die Konstante verwendet:

```golang
type sendingChannel int
const (
  Email sendingChannel = iota
  SMS
  Phone
)
```

In dem Beispiel bekommt Email den Wert 0, SMS den Wert 1 und Phone den Wert 2.

# Kapitel 16: Quiz

Die **Go Proverbs** sind eine Kollektion von guten Ratschlägen von Rob Pike, einem der Erfinder von Go:

- Don't communicate by sharing memory, share memory by communicating. (man sollte eher Channels als Mutexe verwenden, wenn man Daten zwischen Teilen seines Programms teilen will)
- Concurrency is not parallelism.
- Channels orchestrate; mutexes serialize (verwandt mit dem ersten Proverb - Channels steuern den Datenfluss im eigenen Programm, Mutexe führen dazu, dass Dinge hintereinander ausgeführt werden)
- The bigger the interface, the weaker the abstraction.
- Make the zero value useful.
- interface{} says nothing. (da alles dem empty Interface genügt, hat es keinerlei Aussagekraft - ähnlich wie any in Typescript )
- Gofmt's style is no one's favorite, yet gofmt is everyone's favorite. (das ist ein integrierter Bestandteil des Go Ecosystems, darum verwendet es jeder. Und weil es jeder verwendet, schaut jeder Go Code "gleich" aus - auch wenn nicht jeder mit dem Aussehen 100% glücklich ist)
- A little copying is better than a little dependency (Im Gegensatz zu Javascript, wo das node_modules Verzeichnis viel größer als der eigentliche Applikationscode sein kann, ist es bei Go genau umgekehrt - weil nur das, was man unbedingt braucht Teil des Programms wird).
- Syscall must always be guarded with build tags.
- Cgo must always be guarded with build tags.
- Cgo is not Go.
- With the unsafe package there are no guarantees.
- Clear is better than clever (Code wird für Menschen geschrieben, nicht für Maschinen, daher sollte er möglichst klar und "einfach" sein).
- Reflection is never clear.
- Errors are values.
- Don't just check errors, handle them gracefully.
- Design the architecture, name the components, document the details.
- Documentation is for users. (der Code selbst sollte so klar geschrieben sein, dass er sich selbst erklärt. Die Dokumentation sollte eher für die Anwender gedacht sein - wie man das Programm verwendet, nicht für den Maintainer des Programms der den Source Code im Zugriff hat)
- Don't panic.

Auch wenn ich bei weitem nicht alle verstehe ... vielleicht kommt das in weiterer Folge noch.
