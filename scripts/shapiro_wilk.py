"""
Reads two files with one rating per line. Computes the paired
differences and performs Shapiro-Wilk test to assess normality of the difference
distribution.
"""

from scipy.stats import shapiro

def load_and_compute_differences(file_cyber="./cyber.txt", file_traditional="./traditional.txt"):
    with open(file_cyber, 'r') as f_cyber, open(file_traditional, 'r') as f_trad:
        cyber_ratings = [int(line.strip()) for line in f_cyber.readlines()]
        trad_ratings = [int(line.strip()) for line in f_trad.readlines()]
    
    if len(cyber_ratings) != len(trad_ratings):
        raise ValueError("Files have different number of ratings!")
    
    differences = [c - t for c, t in zip(cyber_ratings, trad_ratings)]
    return differences


def main():

    stat, p = shapiro(load_and_compute_differences())

    print(f'Statistic={stat:.4f}, p-value={p:.4f}')

main()
