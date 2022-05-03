# Bachelor Projekt: Dynamic Deadlock Detection in Go
## Personen
- Erik Kassubek, Albert-Ludwigs-Universität Freiburg
- Betreuung: 
    - Prof. Dr. Thiemann, Albert-Ludwigs-Universität Freiburg
    - Prof. Dr. Sulzmann, Hochschule Karlsruhe

## Inhalt
- Theorie über Lockgraphen
- 2-3 Wochen: Analyse von bestehendem Tool zur Deadlock-Detection
    - https://github.com/sasha-s/go-deadlock
- Rest der Zeit: Entwicklung eines eigenen Tools zur Deadlock Detection basierend auf der Theorie
    - Nach Möglichkeit Betrachtung von Channels und read/write locks

## Vorgehen und Ziele


1. Review bestehender Tools zwecks dynamischer Deadlock Analyse in Go.

Betrachte  z.B. https://github.com/sasha-s/go-deadlock.
Gibt es weitere Tools?
Funktionsweise?
Einschränkungen?

2. Umsetzung von UNDEAD in Go.

Ähnlich wie "sasha-s", (RW)Mutex als "drop-in replacements for sync.(RW)Mutex".

3. Evaluation anhand geeigneter Beispiele.

Eigene Beispiele.
Weitere Beispiele siehe http://lujie.ac.cn/files/papers/GoBench.pdf.

4. Erweiterung von UNDEAD mit rw-locks (optional falls noch Zeit vorhanden).