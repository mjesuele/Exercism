(ns hello-world)

(defn hello
  "I say hello"
  [& [name]]
  (if name (str "Hello, " name "!" ) "Hello, World!"))