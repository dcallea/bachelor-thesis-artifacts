"""
Performs a paired sample t-test on two data sets.

Expects two files in the same folder: cyber.txt and traditional.txt

Each file should have one rating per line, and the files should have the
same number of ratings.
"""

from scipy.stats import ttest_rel

file_cyber = "./cyber.txt"
file_traditional = "./traditional.txt"

def main():
    with open(file_cyber, 'r') as f_cyber, open(file_traditional, 'r') as f_trad:
        cyber_ratings = [int(line.strip()) for line in f_cyber.readlines()]
        trad_ratings = [int(line.strip()) for line in f_trad.readlines()]
    
    if len(cyber_ratings) != len(trad_ratings):
        raise ValueError("Files have different number of ratings!")

    stat, p = ttest_rel(cyber_ratings, trad_ratings)

    print(f't-test test statistic={stat}, p-value={p:.4f}')


main()
