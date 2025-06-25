"""
Performs Wilcoxon-signed rank test on dataset, expects one rating per line.
Checks if data is statistically significantly greater than midpoint of 3. 
"""

from scipy.stats import wilcoxon

with open('blend.txt', 'r') as file:
    data = [float(line.strip()) for line in file.readlines()]

differences = [x - 3 for x in data]

stat, p_value = wilcoxon(differences, alternative ='greater')

print(f"Wilcoxon test statistic: {stat}")
print(f"P-value: {p_value}")
