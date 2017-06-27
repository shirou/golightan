syntax highlighter
========================

Syntx highlighter using `ANTLR <http://www.antlr.org/>`_

Formatter is same as `pygments <http://pygments.org/>`_.

.. image:: https://raw.githubusercontent.com/shirou/highlighter/master/terminal_highlight.png
   :alt: teminal highlighter for sqlite3
   :width: 100%
   :align: center


Available Lexers
-----------------

- sqlite3
- golang


lexer is placed at other package. See https://github.com/shirou/antlr-grammars-v4-go



Available Formatters
--------------------


- terminal
  - same as `terminal255`
- raw
- html
  - style is not implemented yet


command line
--------------

::

  go build cmd/highlighter -o highlighter


::

  Usage of ./cmd/highlighter/highlighter:
    -F string
          filters
    -f string
          formatter (default "html")
    -l string
          lexer
    -o string
          output (default "-")
    -s string
          html style (default "default")

LICENSE
============

BSD 3-clause license (same as ANTLR)
