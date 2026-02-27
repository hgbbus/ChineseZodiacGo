# Programming Project

## Generate the 60-Year Chinese Zodiac Cycle

### Overview

In traditional Chinese calendrical systems, years follow a **60-year cycle** formed by combining:

* **10 Heavenly Stems (天干)** → linked to **Yin–Yang** and **Five Elements**
* **12 Earthly Branches (地支)** → linked to **Chinese Zodiac Animals**

Each year is represented by a unique combination of one stem (e.g., 甲) and one branch (e.g., 子), resulting in 60 distinct year names. Details on how the stems and branches combine to form the cycle are provided in later sections.

Because 10 and 12 have a least common multiple of 60, the full pattern repeats every 60 years.

Your task is to build a program that:

1. Generates the full 60-year cycle.
2. Displays the Heavenly Stem, Earthly Branch, Zodiac Animal.
3. Shows Yin or Yang.
4. Shows the associated Five Element.
5. Maps real Gregorian years to the correct cycle name.

## Background Reference

### 1️⃣ Heavenly Stems (10)

| Index | Stem | Pinyin | Element | Yin/Yang |
| ----- | ---- | ------ | ------- | -------- |
| 0     | 甲    | Jia    | Wood    | Yang     |
| 1     | 乙    | Yi     | Wood    | Yin      |
| 2     | 丙    | Bing   | Fire    | Yang     |
| 3     | 丁    | Ding   | Fire    | Yin      |
| 4     | 戊    | Wu     | Earth   | Yang     |
| 5     | 己    | Ji     | Earth   | Yin      |
| 6     | 庚    | Geng   | Metal   | Yang     |
| 7     | 辛    | Xin    | Metal   | Yin      |
| 8     | 壬    | Ren    | Water   | Yang     |
| 9     | 癸    | Gui    | Water   | Yin      |

### 2️⃣ Earthly Branches (12 Zodiac Animals)

| Index | Branch | Animal  |
| ----- | ------ | ------- |
| 0     | 子      | Rat     |
| 1     | 丑      | Ox      |
| 2     | 寅      | Tiger   |
| 3     | 卯      | Rabbit  |
| 4     | 辰      | Dragon  |
| 5     | 巳      | Snake   |
| 6     | 午      | Horse   |
| 7     | 未      | Goat    |
| 8     | 申      | Monkey  |
| 9     | 酉      | Rooster |
| 10    | 戌      | Dog     |
| 11    | 亥      | Pig     |

## How the 60-Year Cycle Works

The first year (both indexes = 0) of the cycle is formed by pairing the first stem (甲) with the first branch (子), resulting in 甲子 (Jia-Zi), which corresponds to the Yang Wood Rat.

For each new year:

* Stem index increases by 1 (mod 10)
* Branch index increases by 1 (mod 12)

After 60 increments, both indexes reset to their starting alignment (both indexes = 0).

Example starting point:

* Year 1984 = 甲子 (Jia-Zi) = Yang Wood Rat
* This is commonly used as the base year.

### Visualization of the Cycle

The following digram visualizes the cycle:
<figure>
  <img src="https://www.thechinastory.org/content/uploads/2021/03/Stems-branches-diagram.jpg" alt="60-Year Cycle Diagram" width="600">
  <figcaption>Figure 1: 60-Year Cycle Formed by Pairing Heavenly Stems and Earthly Branches</figcaption>
</figure>

The above diagram also shows that Year 2020 = 庚子 (Geng-Zi) = Yang Metal Rat.

<br/>

The following diagram may help English speakers visualize the cycle a little better:
<figure>
  <img src="https://ytliu0.github.io/ChineseCalendar/sexagenary.png" alt="60-Year Cycle" width="600">
  <figcaption>Figure 2: 60-Year Cycle Representation in English along with corresponding animals</figcaption>
</figure>

<br/>

# Assignment Requirements

## Part 1 — Data Modeling

Create appropriate data structures for:

* Heavenly Stems (include element + yin/yang)
* Earthly Branches (include animal)

**Hint**: You can simply use lists/arrays for the stems and branches, where the index corresponds to the year offset from the base year (1984). For simplicity, you can separately store the branch names and animal names in _parallel_ lists/arrays, and access them using the same index. You don't need to store the element and yin/yang separately for each year, as they can be derived from the stem index using modular arithmetic (see later sections).

A full year description should include:
  * Gregorian year
  * Stem
  * Branch
  * Animal
  * Element
  * Yin/Yang

---

## Part 2 — Generate the 60-Year Cycle

Write a function:

```python
generate_cycle(start_year)
```

where `start_year` is the Gregorian year corresponding to the first entry in the cycle (e.g., 1984 for Jia-Zi).

It should:

* Generate 60 consecutive years
* Correctly wrap:

  * stem_index % 10
  * branch_index % 12

Output format example (for `start_year` = 1984):
```
1984 – Jia Zi – Yang Wood – Rat
1985 – Yi Chou – Yin Wood – Ox
1986 – Bing Yin – Yang Fire – Tiger
...
```

If the `start_year` is not 1984, the function should still generate the correct cycle starting from that year. For example, if `start_year` is 2020, the output should be:
```
2020 – Geng Zi – Yang Metal – Rat
2021 – Xin Chou – Yin Metal – Ox
2022 – Ren Yin – Yang Water – Tiger
...
```

---

## Part 3 — Year Lookup Function

Write a function:

```python
get_zodiac(year)
```

That:

1. Calculates offset from base year (1984)
2. Uses modular arithmetic to determine:

   * stem_index
   * branch_index
3. Returns full formatted zodiac description

Test cases:

* 1984 → Jia Zi (Yang Wood Rat)
* 2024 → Jia Chen (Yang Wood Dragon)
* 1960 → Geng Zi (Yang Metal Rat)

---

## Part 4 — Yin-Yang & Five Element Logic

Do NOT hardcode element separately for each year.

Instead:

* Determine element from stem index:

  * 0–1 → Wood
  * 2–3 → Fire
  * 4–5 → Earth
  * 6–7 → Metal
  * 8–9 → Water

* Determine Yin/Yang:

  * Even stem index → Yang
  * Odd stem index → Yin

---

## Part 5 - Interface & Output Formatting

Design a clean menu driven interface with menu items for users to:
1. Generate and display the full 60-year cycle starting from a specified year.
2. Look up the zodiac information for a specific Gregorian year.

Ensure the output is well-formatted and easy to read.

# Grading Rubric

| Category                        | Points |
| ------------------------------- | ------ |
| Correct 60-year generation      | 30     |
| Correct modular arithmetic      | 20     |
| Proper element + yin/yang logic | 20     |
| Clean code structure            | 15     |
| Documentation                   | 10     |
| User interface & formatting     | 5      |

Total: 100 points

---

# Expected Learning Outcomes

Students will:

* Practice modular arithmetic
* Use arrays / lists
* Model structured data
* Understand cyclic systems
* Implement real-world calendar logic
* Design clean abstraction layers
* Create simple menu-driven interfaces

(End)