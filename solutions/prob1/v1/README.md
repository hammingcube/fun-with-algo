```sh
cat a.cpp | ./run
```

```json
{"language":"cpp","files":[{"name":"a.cpp","content":"# include \u003ciostream\u003e\n# include \u003cfstream\u003e\n\nusing namespace std;\n\nint func(int a, int b) {\n\treturn a+b;\n}\n\nint main() {\n\tint a, b, result;\n\t//cout \u003c\u003c \"Hello\\n\";\t\n\tifstream ifile(\"STDIN\");\n\twhile(ifile \u003e\u003e a) {\n\t\tifile \u003e\u003e b;\n\t\tresult = func(a,b);\n\t\tcout \u003c\u003c result \u003c\u003c endl;\n\t}\n\treturn 0;\n}\n"},{"name":"STDIN","content":"1 2"}]}
```

```sh
echo '{"language":"cpp","files":[{"name":"a.cpp","content":"# include \u003ciostream\u003e\n# include \u003cfstream\u003e\n\nusing namespace std;\n\nint func(int a, int b) {\n\treturn a+b;\n}\n\nint main() {\n\tint a, b, result;\n\t//cout \u003c\u003c \"Hello\\n\";\t\n\tifstream ifile(\"STDIN\");\n\twhile(ifile \u003e\u003e a) {\n\t\tifile \u003e\u003e b;\n\t\tresult = func(a,b);\n\t\tcout \u003c\u003c result \u003c\u003c endl;\n\t}\n\treturn 0;\n}\n"},{"name":"STDIN","content":"1 2"}]}' | docker run -i glot/clang:latest > output.txt
```


