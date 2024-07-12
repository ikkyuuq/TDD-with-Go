package main;

import "fmt"

const (
  spanish = "Spanish"
  french = "French"

  DefaultHelloPrefix = "Hello, "
  SpanishHelloPrefix = "Hola, "
  FrenchHelloPrefix = "Bonjour, "
)

func Hello(name, lang string) string {
  if name == "" {
    name = "World"
  }
  return greetingPrefix(lang) + name;
}

func greetingPrefix(lang string) (prefix string){
  switch lang {
  case spanish:
    prefix = SpanishHelloPrefix;
  case french:
    prefix = FrenchHelloPrefix;
  default:
    prefix = DefaultHelloPrefix;
  }
  return
}

func main() {
  fmt.Println(Hello("Ikkyuu", ""));
}
