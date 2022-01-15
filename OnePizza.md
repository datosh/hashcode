# One Pizza

[Trial Problem](https://codingcompetitions.withgoogle.com/hashcode/round/00000000008f5ca9/00000000008f6f33)
for HashCode 2022.

## Problem

You are opening a small pizzeria. In fact, your pizzeria is so small that you decided to offer only one type of pizza. Now you need to decide what ingredients to include (peppers? tomatoes? both?).

Everyone has their own pizza preferences. Each of your potential clients has some ingredients they like, and maybe some ingredients they dislike. Each client will come to your pizzeria if both conditions are true:

* all the ingredients they like are on the pizza, and
* none of the ingredients they dislike are on the pizza

Each client is OK with additional ingredients they neither like or dislike being present on the pizza. Your task is to choose which ingredients to put on your only pizza type, to maximize the number of clients that will visit your pizzeria.

## Input

The first line contains one integer 1≤C≤10^5 - the number of potential clients.
The following 2*C lines describe the clients’ preferences in the following format:
* First line contains integer 1≤L≤5, followed by L names of ingredients a client likes, delimited by spaces.
* Second line contains integer 0≤D≤5, followed by D names of ingredients a client dislikes, delimited by spaces.
* Each ingredient name consists of between 1 and 15 ASCII characters. Each character is one of the lowercase letters (a-z) or a digit (0-9).

## Input Data

In folder `one_pizza_input`.

## Submission

The submission should consist of one line consisting of a single number 0≤N followed by a list of N ingredients to put on the only pizza available in the pizzeria, separated by spaces. The list of ingredients should contain only the ingredients mentioned by at least one client, without duplicates.

## Scoring

A solution scores one point for each client that will come to your pizzeria. A client will come to your pizzeria if all the ingredients they like are on the pizza and none of the ingredients they dislike are on the pizza.

## Sample

### Sample Input

```txt
3
2 cheese peppers
0
1 basil
1 pineapple
2 mushrooms tomatoes
1 basil
```

### Sample Output

```
4 cheese mushrooms tomatoes peppers
```