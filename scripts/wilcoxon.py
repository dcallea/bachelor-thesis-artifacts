"""
Performs Wilcoxon signed-rank test. Takes two datasets with same number of
integer ratings (one per line).
"""


from scipy.stats import wilcoxon

file_cyber = "./cyber.txt"
file_traditional = "./traditional.txt"

def main():
    with open(file_cyber, 'r') as f_cyber, open(file_traditional, 'r') as f_trad:
        cyber_ratings = [int(line.strip()) for line in f_cyber.readlines()]
        trad_ratings = [int(line.strip()) for line in f_trad.readlines()]
    
    if len(cyber_ratings) != len(trad_ratings):
        raise ValueError("Files have different number of ratings!")

    stat, p = wilcoxon(cyber_ratings, trad_ratings, alternative='greater')

    print(f'Wilcoxon test statistic={stat}, p-value={p:.4f}')


main()
