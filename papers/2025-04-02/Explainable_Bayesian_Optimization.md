# Explainable Bayesian Optimization

**URL**: http://arxiv.org/abs/2401.13334v2

## 原始摘要

Manual parameter tuning of cyber-physical systems is a common practice, but
it is labor-intensive. Bayesian Optimization (BO) offers an automated
alternative, yet its black-box nature reduces trust and limits human-BO
collaborative system tuning. Experts struggle to interpret BO recommendations
due to the lack of explanations. This paper addresses the post-hoc BO
explainability problem for cyber-physical systems. We introduce TNTRules
(Tune-No-Tune Rules), a novel algorithm that provides both global and local
explanations for BO recommendations. TNTRules generates actionable rules and
visual graphs, identifying optimal solution bounds and ranges, as well as
potential alternative solutions. Unlike existing explainable AI (XAI) methods,
TNTRules is tailored specifically for BO, by encoding uncertainty via a
variance pruning technique and hierarchical agglomerative clustering. A
multi-objective optimization approach allows maximizing explanation quality. We
evaluate TNTRules using established XAI metrics (Correctness, Completeness, and
Compactness) and compare it against adapted baseline methods. The results
demonstrate that TNTRules generates high-fidelity, compact, and complete
explanations, significantly outperforming three baselines on 5 multi-objective
testing functions and 2 hyperparameter tuning problems.


## AI 摘要

本文提出TNTRules算法，解决信息物理系统中贝叶斯优化(BO)的可解释性问题。该算法通过方差剪枝和层次聚类技术，生成全局和局部解释规则及可视化图表，识别最优解范围和替代方案。相比现有方法，TNTRules专为BO设计，采用多目标优化提升解释质量。实验表明，在5个测试函数和2个调参问题上，TNTRules在正确性、完整性和简洁性方面显著优于基线方法，能产生高保真、紧凑且完整的解释，有助于增强用户对BO推荐结果的信任。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-02T12:02:36Z
- **目录日期**: 2025-04-02
