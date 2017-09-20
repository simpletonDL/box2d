package main

// Based on https://github.com/erincatto/Box2D/blob/master/Box2D/Testbed/Tests/CharacterCollision.h

// Note:
// These 4 numbers have been modified from the output of the c++ reference
// These are approximation erros due to the fact that the golang implementation uses float64
// And the C++ implementation uses float32
/*
--- C++
+++ Go
@@ -219 +219 @@
-19(09_circlecharacter1): 3.000 4.504 0.000
+19(09_circlecharacter1): 3.000 4.505 0.000
@@ -416 +416 @@
-37(08_hexagoncharacter): -5.034 6.057 0.000
+37(08_hexagoncharacter): -5.034 6.058 0.000
@@ -438 +438 @@
-39(08_hexagoncharacter): -5.086 6.031 0.000
+39(08_hexagoncharacter): -5.086 6.032 0.000
@@ -471 +471 @@
-42(08_hexagoncharacter): -5.170 5.989 0.000
+42(08_hexagoncharacter): -5.170 5.990 0.000
*/

var expected string = `0(00_ground): 0.000 0.000 0.000
0(01_colinearground): 0.000 0.000 0.000
0(02_chainshape): 0.000 0.000 0.785
0(03_squaretiles): 0.000 0.000 0.000
0(04_edgeloopsquare): 0.000 0.000 0.000
0(05_edgelooppoly): -10.000 4.000 0.000
0(06_squarecharacter1): -3.000 7.997 0.000
0(07_squarecharacter2): -5.000 4.997 0.000
0(08_hexagoncharacter): -5.000 7.997 0.000
0(09_circlecharacter1): 3.000 4.997 0.000
0(10_circlecharacter2): -7.000 5.997 0.000
1(00_ground): 0.000 0.000 0.000
1(01_colinearground): 0.000 0.000 0.000
1(02_chainshape): 0.000 0.000 0.785
1(03_squaretiles): 0.000 0.000 0.000
1(04_edgeloopsquare): 0.000 0.000 0.000
1(05_edgelooppoly): -10.000 4.000 0.000
1(06_squarecharacter1): -3.000 7.992 0.000
1(07_squarecharacter2): -5.000 4.992 0.000
1(08_hexagoncharacter): -5.000 7.992 0.000
1(09_circlecharacter1): 3.000 4.992 0.000
1(10_circlecharacter2): -7.000 5.992 0.000
2(00_ground): 0.000 0.000 0.000
2(01_colinearground): 0.000 0.000 0.000
2(02_chainshape): 0.000 0.000 0.785
2(03_squaretiles): 0.000 0.000 0.000
2(04_edgeloopsquare): 0.000 0.000 0.000
2(05_edgelooppoly): -10.000 4.000 0.000
2(06_squarecharacter1): -3.000 7.983 0.000
2(07_squarecharacter2): -5.000 4.983 0.000
2(08_hexagoncharacter): -5.000 7.983 0.000
2(09_circlecharacter1): 3.000 4.983 0.000
2(10_circlecharacter2): -7.000 5.983 0.000
3(00_ground): 0.000 0.000 0.000
3(01_colinearground): 0.000 0.000 0.000
3(02_chainshape): 0.000 0.000 0.785
3(03_squaretiles): 0.000 0.000 0.000
3(04_edgeloopsquare): 0.000 0.000 0.000
3(05_edgelooppoly): -10.000 4.000 0.000
3(06_squarecharacter1): -3.000 7.972 0.000
3(07_squarecharacter2): -5.000 4.972 0.000
3(08_hexagoncharacter): -5.000 7.972 0.000
3(09_circlecharacter1): 3.000 4.972 0.000
3(10_circlecharacter2): -7.000 5.972 0.000
4(00_ground): 0.000 0.000 0.000
4(01_colinearground): 0.000 0.000 0.000
4(02_chainshape): 0.000 0.000 0.785
4(03_squaretiles): 0.000 0.000 0.000
4(04_edgeloopsquare): 0.000 0.000 0.000
4(05_edgelooppoly): -10.000 4.000 0.000
4(06_squarecharacter1): -3.000 7.958 0.000
4(07_squarecharacter2): -5.000 4.958 0.000
4(08_hexagoncharacter): -5.000 7.958 0.000
4(09_circlecharacter1): 3.000 4.958 0.000
4(10_circlecharacter2): -7.000 5.958 0.000
5(00_ground): 0.000 0.000 0.000
5(01_colinearground): 0.000 0.000 0.000
5(02_chainshape): 0.000 0.000 0.785
5(03_squaretiles): 0.000 0.000 0.000
5(04_edgeloopsquare): 0.000 0.000 0.000
5(05_edgelooppoly): -10.000 4.000 0.000
5(06_squarecharacter1): -3.000 7.942 0.000
5(07_squarecharacter2): -5.000 4.942 0.000
5(08_hexagoncharacter): -5.000 7.942 0.000
5(09_circlecharacter1): 3.000 4.942 0.000
5(10_circlecharacter2): -7.000 5.942 0.000
6(00_ground): 0.000 0.000 0.000
6(01_colinearground): 0.000 0.000 0.000
6(02_chainshape): 0.000 0.000 0.785
6(03_squaretiles): 0.000 0.000 0.000
6(04_edgeloopsquare): 0.000 0.000 0.000
6(05_edgelooppoly): -10.000 4.000 0.000
6(06_squarecharacter1): -3.000 7.922 0.000
6(07_squarecharacter2): -5.000 4.922 0.000
6(08_hexagoncharacter): -5.000 7.922 0.000
6(09_circlecharacter1): 3.000 4.922 0.000
6(10_circlecharacter2): -7.000 5.922 0.000
7(00_ground): 0.000 0.000 0.000
7(01_colinearground): 0.000 0.000 0.000
7(02_chainshape): 0.000 0.000 0.785
7(03_squaretiles): 0.000 0.000 0.000
7(04_edgeloopsquare): 0.000 0.000 0.000
7(05_edgelooppoly): -10.000 4.000 0.000
7(06_squarecharacter1): -3.000 7.900 0.000
7(07_squarecharacter2): -5.000 4.900 0.000
7(08_hexagoncharacter): -5.000 7.900 0.000
7(09_circlecharacter1): 3.000 4.900 0.000
7(10_circlecharacter2): -7.000 5.900 0.000
8(00_ground): 0.000 0.000 0.000
8(01_colinearground): 0.000 0.000 0.000
8(02_chainshape): 0.000 0.000 0.785
8(03_squaretiles): 0.000 0.000 0.000
8(04_edgeloopsquare): 0.000 0.000 0.000
8(05_edgelooppoly): -10.000 4.000 0.000
8(06_squarecharacter1): -3.000 7.875 0.000
8(07_squarecharacter2): -5.000 4.875 0.000
8(08_hexagoncharacter): -5.000 7.875 0.000
8(09_circlecharacter1): 3.000 4.875 0.000
8(10_circlecharacter2): -7.000 5.875 0.000
9(00_ground): 0.000 0.000 0.000
9(01_colinearground): 0.000 0.000 0.000
9(02_chainshape): 0.000 0.000 0.785
9(03_squaretiles): 0.000 0.000 0.000
9(04_edgeloopsquare): 0.000 0.000 0.000
9(05_edgelooppoly): -10.000 4.000 0.000
9(06_squarecharacter1): -3.000 7.847 0.000
9(07_squarecharacter2): -5.000 4.847 0.000
9(08_hexagoncharacter): -5.000 7.847 0.000
9(09_circlecharacter1): 3.000 4.847 0.000
9(10_circlecharacter2): -7.000 5.847 0.000
10(00_ground): 0.000 0.000 0.000
10(01_colinearground): 0.000 0.000 0.000
10(02_chainshape): 0.000 0.000 0.785
10(03_squaretiles): 0.000 0.000 0.000
10(04_edgeloopsquare): 0.000 0.000 0.000
10(05_edgelooppoly): -10.000 4.000 0.000
10(06_squarecharacter1): -3.000 7.817 0.000
10(07_squarecharacter2): -5.000 4.817 0.000
10(08_hexagoncharacter): -5.000 7.817 0.000
10(09_circlecharacter1): 3.000 4.817 0.000
10(10_circlecharacter2): -7.000 5.817 0.000
11(00_ground): 0.000 0.000 0.000
11(01_colinearground): 0.000 0.000 0.000
11(02_chainshape): 0.000 0.000 0.785
11(03_squaretiles): 0.000 0.000 0.000
11(04_edgeloopsquare): 0.000 0.000 0.000
11(05_edgelooppoly): -10.000 4.000 0.000
11(06_squarecharacter1): -3.000 7.783 0.000
11(07_squarecharacter2): -5.000 4.783 0.000
11(08_hexagoncharacter): -5.000 7.783 0.000
11(09_circlecharacter1): 3.000 4.783 0.000
11(10_circlecharacter2): -7.000 5.783 0.000
12(00_ground): 0.000 0.000 0.000
12(01_colinearground): 0.000 0.000 0.000
12(02_chainshape): 0.000 0.000 0.785
12(03_squaretiles): 0.000 0.000 0.000
12(04_edgeloopsquare): 0.000 0.000 0.000
12(05_edgelooppoly): -10.000 4.000 0.000
12(06_squarecharacter1): -3.000 7.747 0.000
12(07_squarecharacter2): -5.000 4.747 0.000
12(08_hexagoncharacter): -5.000 7.747 0.000
12(09_circlecharacter1): 3.000 4.747 0.000
12(10_circlecharacter2): -6.990 5.779 -0.043
13(00_ground): 0.000 0.000 0.000
13(01_colinearground): 0.000 0.000 0.000
13(02_chainshape): 0.000 0.000 0.785
13(03_squaretiles): 0.000 0.000 0.000
13(04_edgeloopsquare): 0.000 0.000 0.000
13(05_edgelooppoly): -10.000 4.000 0.000
13(06_squarecharacter1): -3.000 7.708 0.000
13(07_squarecharacter2): -5.000 4.708 0.000
13(08_hexagoncharacter): -5.000 7.708 0.000
13(09_circlecharacter1): 3.000 4.708 0.000
13(10_circlecharacter2): -6.980 5.774 -0.090
14(00_ground): 0.000 0.000 0.000
14(01_colinearground): 0.000 0.000 0.000
14(02_chainshape): 0.000 0.000 0.785
14(03_squaretiles): 0.000 0.000 0.000
14(04_edgeloopsquare): 0.000 0.000 0.000
14(05_edgelooppoly): -10.000 4.000 0.000
14(06_squarecharacter1): -3.000 7.667 0.000
14(07_squarecharacter2): -5.000 4.667 0.000
14(08_hexagoncharacter): -5.000 7.667 0.000
14(09_circlecharacter1): 3.000 4.667 0.000
14(10_circlecharacter2): -6.969 5.769 -0.140
15(00_ground): 0.000 0.000 0.000
15(01_colinearground): 0.000 0.000 0.000
15(02_chainshape): 0.000 0.000 0.785
15(03_squaretiles): 0.000 0.000 0.000
15(04_edgeloopsquare): 0.000 0.000 0.000
15(05_edgelooppoly): -10.000 4.000 0.000
15(06_squarecharacter1): -3.000 7.622 0.000
15(07_squarecharacter2): -5.000 4.622 0.000
15(08_hexagoncharacter): -5.000 7.622 0.000
15(09_circlecharacter1): 3.000 4.622 0.000
15(10_circlecharacter2): -6.957 5.763 -0.193
16(00_ground): 0.000 0.000 0.000
16(01_colinearground): 0.000 0.000 0.000
16(02_chainshape): 0.000 0.000 0.785
16(03_squaretiles): 0.000 0.000 0.000
16(04_edgeloopsquare): 0.000 0.000 0.000
16(05_edgelooppoly): -10.000 4.000 0.000
16(06_squarecharacter1): -3.000 7.575 0.000
16(07_squarecharacter2): -5.000 4.575 0.000
16(08_hexagoncharacter): -5.000 7.575 0.000
16(09_circlecharacter1): 3.000 4.575 0.000
16(10_circlecharacter2): -6.944 5.757 -0.249
17(00_ground): 0.000 0.000 0.000
17(01_colinearground): 0.000 0.000 0.000
17(02_chainshape): 0.000 0.000 0.785
17(03_squaretiles): 0.000 0.000 0.000
17(04_edgeloopsquare): 0.000 0.000 0.000
17(05_edgelooppoly): -10.000 4.000 0.000
17(06_squarecharacter1): -3.000 7.525 0.000
17(07_squarecharacter2): -5.000 4.525 0.000
17(08_hexagoncharacter): -5.000 7.525 0.000
17(09_circlecharacter1): 3.000 4.525 0.000
17(10_circlecharacter2): -6.931 5.750 -0.309
18(00_ground): 0.000 0.000 0.000
18(01_colinearground): 0.000 0.000 0.000
18(02_chainshape): 0.000 0.000 0.785
18(03_squaretiles): 0.000 0.000 0.000
18(04_edgeloopsquare): 0.000 0.000 0.000
18(05_edgelooppoly): -10.000 4.000 0.000
18(06_squarecharacter1): -3.000 7.472 0.000
18(07_squarecharacter2): -5.000 4.472 0.000
18(08_hexagoncharacter): -5.000 7.472 0.000
18(09_circlecharacter1): 3.000 4.504 0.000
18(10_circlecharacter2): -6.917 5.743 -0.372
19(00_ground): 0.000 0.000 0.000
19(01_colinearground): 0.000 0.000 0.000
19(02_chainshape): 0.000 0.000 0.785
19(03_squaretiles): 0.000 0.000 0.000
19(04_edgeloopsquare): 0.000 0.000 0.000
19(05_edgelooppoly): -10.000 4.000 0.000
19(06_squarecharacter1): -3.000 7.417 0.000
19(07_squarecharacter2): -5.000 4.417 0.000
19(08_hexagoncharacter): -5.000 7.417 0.000
19(09_circlecharacter1): 3.000 4.505 0.000
19(10_circlecharacter2): -6.902 5.736 -0.439
20(00_ground): 0.000 0.000 0.000
20(01_colinearground): 0.000 0.000 0.000
20(02_chainshape): 0.000 0.000 0.785
20(03_squaretiles): 0.000 0.000 0.000
20(04_edgeloopsquare): 0.000 0.000 0.000
20(05_edgelooppoly): -10.000 4.000 0.000
20(06_squarecharacter1): -3.000 7.358 0.000
20(07_squarecharacter2): -5.000 4.358 0.000
20(08_hexagoncharacter): -5.000 7.358 0.000
20(09_circlecharacter1): 3.000 4.505 0.000
20(10_circlecharacter2): -6.887 5.728 -0.509
21(00_ground): 0.000 0.000 0.000
21(01_colinearground): 0.000 0.000 0.000
21(02_chainshape): 0.000 0.000 0.785
21(03_squaretiles): 0.000 0.000 0.000
21(04_edgeloopsquare): 0.000 0.000 0.000
21(05_edgelooppoly): -10.000 4.000 0.000
21(06_squarecharacter1): -3.000 7.297 0.000
21(07_squarecharacter2): -5.000 4.297 0.000
21(08_hexagoncharacter): -5.000 7.297 0.000
21(09_circlecharacter1): 3.000 4.505 0.000
21(10_circlecharacter2): -6.871 5.720 -0.582
22(00_ground): 0.000 0.000 0.000
22(01_colinearground): 0.000 0.000 0.000
22(02_chainshape): 0.000 0.000 0.785
22(03_squaretiles): 0.000 0.000 0.000
22(04_edgeloopsquare): 0.000 0.000 0.000
22(05_edgelooppoly): -10.000 4.000 0.000
22(06_squarecharacter1): -3.000 7.233 0.000
22(07_squarecharacter2): -5.000 4.265 0.000
22(08_hexagoncharacter): -5.000 7.233 0.000
22(09_circlecharacter1): 3.000 4.505 0.000
22(10_circlecharacter2): -6.854 5.712 -0.658
23(00_ground): 0.000 0.000 0.000
23(01_colinearground): 0.000 0.000 0.000
23(02_chainshape): 0.000 0.000 0.785
23(03_squaretiles): 0.000 0.000 0.000
23(04_edgeloopsquare): 0.000 0.000 0.000
23(05_edgelooppoly): -10.000 4.000 0.000
23(06_squarecharacter1): -3.000 7.167 0.000
23(07_squarecharacter2): -5.000 4.265 0.000
23(08_hexagoncharacter): -5.000 7.167 0.000
23(09_circlecharacter1): 3.000 4.505 0.000
23(10_circlecharacter2): -6.836 5.703 -0.738
24(00_ground): 0.000 0.000 0.000
24(01_colinearground): 0.000 0.000 0.000
24(02_chainshape): 0.000 0.000 0.785
24(03_squaretiles): 0.000 0.000 0.000
24(04_edgeloopsquare): 0.000 0.000 0.000
24(05_edgelooppoly): -10.000 4.000 0.000
24(06_squarecharacter1): -3.000 7.097 0.000
24(07_squarecharacter2): -5.000 4.265 0.000
24(08_hexagoncharacter): -5.000 7.097 0.000
24(09_circlecharacter1): 3.000 4.505 0.000
24(10_circlecharacter2): -6.818 5.694 -0.821
25(00_ground): 0.000 0.000 0.000
25(01_colinearground): 0.000 0.000 0.000
25(02_chainshape): 0.000 0.000 0.785
25(03_squaretiles): 0.000 0.000 0.000
25(04_edgeloopsquare): 0.000 0.000 0.000
25(05_edgelooppoly): -10.000 4.000 0.000
25(06_squarecharacter1): -3.000 7.025 0.000
25(07_squarecharacter2): -5.000 4.265 0.000
25(08_hexagoncharacter): -5.000 7.025 0.000
25(09_circlecharacter1): 3.000 4.505 0.000
25(10_circlecharacter2): -6.799 5.684 -0.907
26(00_ground): 0.000 0.000 0.000
26(01_colinearground): 0.000 0.000 0.000
26(02_chainshape): 0.000 0.000 0.785
26(03_squaretiles): 0.000 0.000 0.000
26(04_edgeloopsquare): 0.000 0.000 0.000
26(05_edgelooppoly): -10.000 4.000 0.000
26(06_squarecharacter1): -3.000 6.950 0.000
26(07_squarecharacter2): -5.000 4.265 0.000
26(08_hexagoncharacter): -5.000 6.950 0.000
26(09_circlecharacter1): 3.000 4.505 0.000
26(10_circlecharacter2): -6.779 5.674 -0.997
27(00_ground): 0.000 0.000 0.000
27(01_colinearground): 0.000 0.000 0.000
27(02_chainshape): 0.000 0.000 0.785
27(03_squaretiles): 0.000 0.000 0.000
27(04_edgeloopsquare): 0.000 0.000 0.000
27(05_edgelooppoly): -10.000 4.000 0.000
27(06_squarecharacter1): -3.000 6.872 0.000
27(07_squarecharacter2): -5.000 4.265 0.000
27(08_hexagoncharacter): -5.000 6.872 0.000
27(09_circlecharacter1): 3.000 4.505 0.000
27(10_circlecharacter2): -6.758 5.664 -1.090
28(00_ground): 0.000 0.000 0.000
28(01_colinearground): 0.000 0.000 0.000
28(02_chainshape): 0.000 0.000 0.785
28(03_squaretiles): 0.000 0.000 0.000
28(04_edgeloopsquare): 0.000 0.000 0.000
28(05_edgelooppoly): -10.000 4.000 0.000
28(06_squarecharacter1): -3.000 6.792 0.000
28(07_squarecharacter2): -5.000 4.265 0.000
28(08_hexagoncharacter): -5.000 6.792 0.000
28(09_circlecharacter1): 3.000 4.505 0.000
28(10_circlecharacter2): -6.737 5.654 -1.186
29(00_ground): 0.000 0.000 0.000
29(01_colinearground): 0.000 0.000 0.000
29(02_chainshape): 0.000 0.000 0.785
29(03_squaretiles): 0.000 0.000 0.000
29(04_edgeloopsquare): 0.000 0.000 0.000
29(05_edgelooppoly): -10.000 4.000 0.000
29(06_squarecharacter1): -3.000 6.708 0.000
29(07_squarecharacter2): -5.000 4.265 0.000
29(08_hexagoncharacter): -5.000 6.708 0.000
29(09_circlecharacter1): 3.000 4.505 0.000
29(10_circlecharacter2): -6.715 5.642 -1.286
30(00_ground): 0.000 0.000 0.000
30(01_colinearground): 0.000 0.000 0.000
30(02_chainshape): 0.000 0.000 0.785
30(03_squaretiles): 0.000 0.000 0.000
30(04_edgeloopsquare): 0.000 0.000 0.000
30(05_edgelooppoly): -10.000 4.000 0.000
30(06_squarecharacter1): -3.000 6.622 0.000
30(07_squarecharacter2): -5.000 4.265 0.000
30(08_hexagoncharacter): -5.000 6.622 0.000
30(09_circlecharacter1): 3.000 4.505 0.000
30(10_circlecharacter2): -6.692 5.631 -1.389
31(00_ground): 0.000 0.000 0.000
31(01_colinearground): 0.000 0.000 0.000
31(02_chainshape): 0.000 0.000 0.785
31(03_squaretiles): 0.000 0.000 0.000
31(04_edgeloopsquare): 0.000 0.000 0.000
31(05_edgelooppoly): -10.000 4.000 0.000
31(06_squarecharacter1): -3.000 6.533 0.000
31(07_squarecharacter2): -5.000 4.265 0.000
31(08_hexagoncharacter): -5.000 6.533 0.000
31(09_circlecharacter1): 3.000 4.505 0.000
31(10_circlecharacter2): -6.669 5.619 -1.495
32(00_ground): 0.000 0.000 0.000
32(01_colinearground): 0.000 0.000 0.000
32(02_chainshape): 0.000 0.000 0.785
32(03_squaretiles): 0.000 0.000 0.000
32(04_edgeloopsquare): 0.000 0.000 0.000
32(05_edgelooppoly): -10.000 4.000 0.000
32(06_squarecharacter1): -3.000 6.442 0.000
32(07_squarecharacter2): -5.000 4.265 0.000
32(08_hexagoncharacter): -5.000 6.442 0.000
32(09_circlecharacter1): 3.000 4.505 0.000
32(10_circlecharacter2): -6.644 5.607 -1.605
33(00_ground): 0.000 0.000 0.000
33(01_colinearground): 0.000 0.000 0.000
33(02_chainshape): 0.000 0.000 0.785
33(03_squaretiles): 0.000 0.000 0.000
33(04_edgeloopsquare): 0.000 0.000 0.000
33(05_edgelooppoly): -10.000 4.000 0.000
33(06_squarecharacter1): -3.000 6.347 0.000
33(07_squarecharacter2): -5.000 4.265 0.000
33(08_hexagoncharacter): -5.000 6.347 0.000
33(09_circlecharacter1): 3.000 4.505 0.000
33(10_circlecharacter2): -6.619 5.595 -1.718
34(00_ground): 0.000 0.000 0.000
34(01_colinearground): 0.000 0.000 0.000
34(02_chainshape): 0.000 0.000 0.785
34(03_squaretiles): 0.000 0.000 0.000
34(04_edgeloopsquare): 0.000 0.000 0.000
34(05_edgelooppoly): -10.000 4.000 0.000
34(06_squarecharacter1): -3.000 6.250 0.000
34(07_squarecharacter2): -5.000 4.265 0.000
34(08_hexagoncharacter): -5.000 6.250 0.000
34(09_circlecharacter1): 3.000 4.505 0.000
34(10_circlecharacter2): -6.594 5.582 -1.834
35(00_ground): 0.000 0.000 0.000
35(01_colinearground): 0.000 0.000 0.000
35(02_chainshape): 0.000 0.000 0.785
35(03_squaretiles): 0.000 0.000 0.000
35(04_edgeloopsquare): 0.000 0.000 0.000
35(05_edgelooppoly): -10.000 4.000 0.000
35(06_squarecharacter1): -3.000 6.150 0.000
35(07_squarecharacter2): -5.000 4.265 0.000
35(08_hexagoncharacter): -5.000 6.150 0.000
35(09_circlecharacter1): 3.000 4.505 0.000
35(10_circlecharacter2): -6.567 5.569 -1.954
36(00_ground): 0.000 0.000 0.000
36(01_colinearground): 0.000 0.000 0.000
36(02_chainshape): 0.000 0.000 0.785
36(03_squaretiles): 0.000 0.000 0.000
36(04_edgeloopsquare): 0.000 0.000 0.000
36(05_edgelooppoly): -10.000 4.000 0.000
36(06_squarecharacter1): -3.000 6.047 0.000
36(07_squarecharacter2): -5.000 4.265 0.000
36(08_hexagoncharacter): -5.008 6.070 0.000
36(09_circlecharacter1): 3.000 4.505 0.000
36(10_circlecharacter2): -6.540 5.555 -2.076
37(00_ground): 0.000 0.000 0.000
37(01_colinearground): 0.000 0.000 0.000
37(02_chainshape): 0.000 0.000 0.785
37(03_squaretiles): 0.000 0.000 0.000
37(04_edgeloopsquare): 0.000 0.000 0.000
37(05_edgelooppoly): -10.000 4.000 0.000
37(06_squarecharacter1): -3.000 5.942 0.000
37(07_squarecharacter2): -5.000 4.265 0.000
37(08_hexagoncharacter): -5.034 6.058 0.000
37(09_circlecharacter1): 3.000 4.505 0.000
37(10_circlecharacter2): -6.512 5.541 -2.203
38(00_ground): 0.000 0.000 0.000
38(01_colinearground): 0.000 0.000 0.000
38(02_chainshape): 0.000 0.000 0.785
38(03_squaretiles): 0.000 0.000 0.000
38(04_edgeloopsquare): 0.000 0.000 0.000
38(05_edgelooppoly): -10.000 4.000 0.000
38(06_squarecharacter1): -3.000 5.833 0.000
38(07_squarecharacter2): -5.000 4.265 0.000
38(08_hexagoncharacter): -5.060 6.045 0.000
38(09_circlecharacter1): 3.000 4.505 0.000
38(10_circlecharacter2): -6.483 5.527 -2.332
39(00_ground): 0.000 0.000 0.000
39(01_colinearground): 0.000 0.000 0.000
39(02_chainshape): 0.000 0.000 0.785
39(03_squaretiles): 0.000 0.000 0.000
39(04_edgeloopsquare): 0.000 0.000 0.000
39(05_edgelooppoly): -10.000 4.000 0.000
39(06_squarecharacter1): -3.000 5.722 0.000
39(07_squarecharacter2): -5.000 4.265 0.000
39(08_hexagoncharacter): -5.086 6.032 0.000
39(09_circlecharacter1): 3.000 4.505 0.000
39(10_circlecharacter2): -6.454 5.512 -2.465
40(00_ground): 0.000 0.000 0.000
40(01_colinearground): 0.000 0.000 0.000
40(02_chainshape): 0.000 0.000 0.785
40(03_squaretiles): 0.000 0.000 0.000
40(04_edgeloopsquare): 0.000 0.000 0.000
40(05_edgelooppoly): -10.000 4.000 0.000
40(06_squarecharacter1): -3.000 5.608 0.000
40(07_squarecharacter2): -5.000 4.265 0.000
40(08_hexagoncharacter): -5.114 6.018 0.000
40(09_circlecharacter1): 3.000 4.505 0.000
40(10_circlecharacter2): -6.424 5.497 -2.601
41(00_ground): 0.000 0.000 0.000
41(01_colinearground): 0.000 0.000 0.000
41(02_chainshape): 0.000 0.000 0.785
41(03_squaretiles): 0.000 0.000 0.000
41(04_edgeloopsquare): 0.000 0.000 0.000
41(05_edgelooppoly): -10.000 4.000 0.000
41(06_squarecharacter1): -3.000 5.492 0.000
41(07_squarecharacter2): -5.000 4.265 0.000
41(08_hexagoncharacter): -5.142 6.004 0.000
41(09_circlecharacter1): 3.000 4.505 0.000
41(10_circlecharacter2): -6.393 5.481 -2.741
42(00_ground): 0.000 0.000 0.000
42(01_colinearground): 0.000 0.000 0.000
42(02_chainshape): 0.000 0.000 0.785
42(03_squaretiles): 0.000 0.000 0.000
42(04_edgeloopsquare): 0.000 0.000 0.000
42(05_edgelooppoly): -10.000 4.000 0.000
42(06_squarecharacter1): -3.000 5.372 0.000
42(07_squarecharacter2): -5.000 4.265 0.000
42(08_hexagoncharacter): -5.170 5.990 0.000
42(09_circlecharacter1): 3.000 4.505 0.000
42(10_circlecharacter2): -6.361 5.466 -2.884
43(00_ground): 0.000 0.000 0.000
43(01_colinearground): 0.000 0.000 0.000
43(02_chainshape): 0.000 0.000 0.785
43(03_squaretiles): 0.000 0.000 0.000
43(04_edgeloopsquare): 0.000 0.000 0.000
43(05_edgelooppoly): -10.000 4.000 0.000
43(06_squarecharacter1): -3.000 5.250 0.000
43(07_squarecharacter2): -5.000 4.265 0.000
43(08_hexagoncharacter): -5.200 5.975 0.000
43(09_circlecharacter1): 3.000 4.505 0.000
43(10_circlecharacter2): -6.329 5.449 -3.030
44(00_ground): 0.000 0.000 0.000
44(01_colinearground): 0.000 0.000 0.000
44(02_chainshape): 0.000 0.000 0.785
44(03_squaretiles): 0.000 0.000 0.000
44(04_edgeloopsquare): 0.000 0.000 0.000
44(05_edgelooppoly): -10.000 4.000 0.000
44(06_squarecharacter1): -3.000 5.125 0.000
44(07_squarecharacter2): -5.000 4.265 0.000
44(08_hexagoncharacter): -5.230 5.960 0.000
44(09_circlecharacter1): 3.000 4.505 0.000
44(10_circlecharacter2): -6.296 5.433 -3.179
45(00_ground): 0.000 0.000 0.000
45(01_colinearground): 0.000 0.000 0.000
45(02_chainshape): 0.000 0.000 0.785
45(03_squaretiles): 0.000 0.000 0.000
45(04_edgeloopsquare): 0.000 0.000 0.000
45(05_edgelooppoly): -10.000 4.000 0.000
45(06_squarecharacter1): -3.000 4.997 0.000
45(07_squarecharacter2): -5.000 4.265 0.000
45(08_hexagoncharacter): -5.260 5.945 0.000
45(09_circlecharacter1): 3.000 4.505 0.000
45(10_circlecharacter2): -6.262 5.416 -3.332
46(00_ground): 0.000 0.000 0.000
46(01_colinearground): 0.000 0.000 0.000
46(02_chainshape): 0.000 0.000 0.785
46(03_squaretiles): 0.000 0.000 0.000
46(04_edgeloopsquare): 0.000 0.000 0.000
46(05_edgelooppoly): -10.000 4.000 0.000
46(06_squarecharacter1): -3.000 4.867 0.000
46(07_squarecharacter2): -5.000 4.265 0.000
46(08_hexagoncharacter): -5.292 5.929 0.000
46(09_circlecharacter1): 3.000 4.505 0.000
46(10_circlecharacter2): -6.227 5.399 -3.488
47(00_ground): 0.000 0.000 0.000
47(01_colinearground): 0.000 0.000 0.000
47(02_chainshape): 0.000 0.000 0.785
47(03_squaretiles): 0.000 0.000 0.000
47(04_edgeloopsquare): 0.000 0.000 0.000
47(05_edgelooppoly): -10.000 4.000 0.000
47(06_squarecharacter1): -3.000 4.733 0.000
47(07_squarecharacter2): -5.000 4.265 0.000
47(08_hexagoncharacter): -5.324 5.913 0.000
47(09_circlecharacter1): 3.000 4.505 0.000
47(10_circlecharacter2): -6.192 5.381 -3.648
48(00_ground): 0.000 0.000 0.000
48(01_colinearground): 0.000 0.000 0.000
48(02_chainshape): 0.000 0.000 0.785
48(03_squaretiles): 0.000 0.000 0.000
48(04_edgeloopsquare): 0.000 0.000 0.000
48(05_edgelooppoly): -10.000 4.000 0.000
48(06_squarecharacter1): -3.000 4.597 0.000
48(07_squarecharacter2): -5.000 4.265 0.000
48(08_hexagoncharacter): -5.356 5.897 0.000
48(09_circlecharacter1): 3.000 4.505 0.000
48(10_circlecharacter2): -6.156 5.363 -3.811
49(00_ground): 0.000 0.000 0.000
49(01_colinearground): 0.000 0.000 0.000
49(02_chainshape): 0.000 0.000 0.785
49(03_squaretiles): 0.000 0.000 0.000
49(04_edgeloopsquare): 0.000 0.000 0.000
49(05_edgelooppoly): -10.000 4.000 0.000
49(06_squarecharacter1): -3.000 4.458 0.000
49(07_squarecharacter2): -5.000 4.265 0.000
49(08_hexagoncharacter): -5.390 5.880 0.000
49(09_circlecharacter1): 3.000 4.505 0.000
49(10_circlecharacter2): -6.119 5.345 -3.977
50(00_ground): 0.000 0.000 0.000
50(01_colinearground): 0.000 0.000 0.000
50(02_chainshape): 0.000 0.000 0.785
50(03_squaretiles): 0.000 0.000 0.000
50(04_edgeloopsquare): 0.000 0.000 0.000
50(05_edgelooppoly): -10.000 4.000 0.000
50(06_squarecharacter1): -3.000 4.317 0.000
50(07_squarecharacter2): -5.000 4.265 0.000
50(08_hexagoncharacter): -5.424 5.863 0.000
50(09_circlecharacter1): 3.000 4.505 0.000
50(10_circlecharacter2): -6.082 5.326 -4.146
51(00_ground): 0.000 0.000 0.000
51(01_colinearground): 0.000 0.000 0.000
51(02_chainshape): 0.000 0.000 0.785
51(03_squaretiles): 0.000 0.000 0.000
51(04_edgeloopsquare): 0.000 0.000 0.000
51(05_edgelooppoly): -10.000 4.000 0.000
51(06_squarecharacter1): -3.000 4.172 0.000
51(07_squarecharacter2): -5.000 4.265 0.000
51(08_hexagoncharacter): -5.458 5.846 0.000
51(09_circlecharacter1): 3.000 4.505 0.000
51(10_circlecharacter2): -6.043 5.307 -4.319
52(00_ground): 0.000 0.000 0.000
52(01_colinearground): 0.000 0.000 0.000
52(02_chainshape): 0.000 0.000 0.785
52(03_squaretiles): 0.000 0.000 0.000
52(04_edgeloopsquare): 0.000 0.000 0.000
52(05_edgelooppoly): -10.000 4.000 0.000
52(06_squarecharacter1): -3.000 4.025 0.000
52(07_squarecharacter2): -5.000 4.265 0.000
52(08_hexagoncharacter): -5.494 5.828 0.000
52(09_circlecharacter1): 3.000 4.505 0.000
52(10_circlecharacter2): -6.004 5.287 -4.495
53(00_ground): 0.000 0.000 0.000
53(01_colinearground): 0.000 0.000 0.000
53(02_chainshape): 0.000 0.000 0.785
53(03_squaretiles): 0.000 0.000 0.000
53(04_edgeloopsquare): 0.000 0.000 0.000
53(05_edgelooppoly): -10.000 4.000 0.000
53(06_squarecharacter1): -3.000 3.875 0.000
53(07_squarecharacter2): -5.000 4.265 0.000
53(08_hexagoncharacter): -5.530 5.810 0.000
53(09_circlecharacter1): 3.000 4.505 0.000
53(10_circlecharacter2): -5.976 5.301 -4.620
54(00_ground): 0.000 0.000 0.000
54(01_colinearground): 0.000 0.000 0.000
54(02_chainshape): 0.000 0.000 0.785
54(03_squaretiles): 0.000 0.000 0.000
54(04_edgeloopsquare): 0.000 0.000 0.000
54(05_edgelooppoly): -10.000 4.000 0.000
54(06_squarecharacter1): -3.000 3.722 0.000
54(07_squarecharacter2): -5.000 4.265 0.000
54(08_hexagoncharacter): -5.547 5.802 0.000
54(09_circlecharacter1): 3.000 4.505 0.000
54(10_circlecharacter2): -6.017 5.288 -4.653
55(00_ground): 0.000 0.000 0.000
55(01_colinearground): 0.000 0.000 0.000
55(02_chainshape): 0.000 0.000 0.785
55(03_squaretiles): 0.000 0.000 0.000
55(04_edgeloopsquare): 0.000 0.000 0.000
55(05_edgelooppoly): -10.000 4.000 0.000
55(06_squarecharacter1): -3.000 3.567 0.000
55(07_squarecharacter2): -5.000 4.265 0.000
55(08_hexagoncharacter): -5.554 5.804 0.000
55(09_circlecharacter1): 3.000 4.505 0.000
55(10_circlecharacter2): -6.029 5.290 -4.630
56(00_ground): 0.000 0.000 0.000
56(01_colinearground): 0.000 0.000 0.000
56(02_chainshape): 0.000 0.000 0.785
56(03_squaretiles): 0.000 0.000 0.000
56(04_edgeloopsquare): 0.000 0.000 0.000
56(05_edgelooppoly): -10.000 4.000 0.000
56(06_squarecharacter1): -3.000 3.408 0.000
56(07_squarecharacter2): -5.000 4.265 0.000
56(08_hexagoncharacter): -5.555 5.811 0.000
56(09_circlecharacter1): 3.000 4.505 0.000
56(10_circlecharacter2): -6.035 5.293 -4.619
57(00_ground): 0.000 0.000 0.000
57(01_colinearground): 0.000 0.000 0.000
57(02_chainshape): 0.000 0.000 0.785
57(03_squaretiles): 0.000 0.000 0.000
57(04_edgeloopsquare): 0.000 0.000 0.000
57(05_edgelooppoly): -10.000 4.000 0.000
57(06_squarecharacter1): -3.000 3.247 0.000
57(07_squarecharacter2): -5.000 4.265 0.000
57(08_hexagoncharacter): -5.555 5.816 0.000
57(09_circlecharacter1): 3.000 4.505 0.000
57(10_circlecharacter2): -6.038 5.297 -4.612
58(00_ground): 0.000 0.000 0.000
58(01_colinearground): 0.000 0.000 0.000
58(02_chainshape): 0.000 0.000 0.785
58(03_squaretiles): 0.000 0.000 0.000
58(04_edgeloopsquare): 0.000 0.000 0.000
58(05_edgelooppoly): -10.000 4.000 0.000
58(06_squarecharacter1): -3.000 3.083 0.000
58(07_squarecharacter2): -5.000 4.265 0.000
58(08_hexagoncharacter): -5.556 5.818 0.000
58(09_circlecharacter1): 3.000 4.505 0.000
58(10_circlecharacter2): -6.040 5.298 -4.607
59(00_ground): 0.000 0.000 0.000
59(01_colinearground): 0.000 0.000 0.000
59(02_chainshape): 0.000 0.000 0.785
59(03_squaretiles): 0.000 0.000 0.000
59(04_edgeloopsquare): 0.000 0.000 0.000
59(05_edgelooppoly): -10.000 4.000 0.000
59(06_squarecharacter1): -3.000 2.917 0.000
59(07_squarecharacter2): -5.000 4.265 0.000
59(08_hexagoncharacter): -5.556 5.818 0.000
59(09_circlecharacter1): 3.000 4.505 0.000
59(10_circlecharacter2): -6.041 5.299 -4.606
`
