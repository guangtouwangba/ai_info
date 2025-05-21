# Semi-parametric efficient estimation of small genetic effects in large-scale population cohorts

**URL**: http://arxiv.org/abs/2505.14675v1

## 原始摘要

Population genetics seeks to quantify DNA variant associations with traits or
diseases, as well as interactions among variants and with environmental
factors. Computing millions of estimates in large cohorts in which small effect
sizes are expected, necessitates minimising model-misspecification bias to
control false discoveries. We present TarGene, a unified statistical workflow
for the semi-parametric efficient and double robust estimation of genetic
effects including k-point interactions among categorical variables in the
presence of confounding and weak population dependence. k-point interactions,
or Average Interaction Effects (AIEs), are a direct generalisation of the usual
average treatment effect (ATE). We estimate AIEs with cross-validated and/or
weighted versions of Targeted Minimum Loss-based Estimators (TMLE) and One-Step
Estimators (OSE). The effect of dependence among data units on variance
estimates is corrected by using sieve plateau variance estimators based on
genetic relatedness across the units. We present extensive realistic
simulations to demonstrate power, coverage, and control of type I error. Our
motivating application is the targeted estimation of genetic effects on trait,
including two-point and higher-order gene-gene and gene-environment
interactions, in large-scale genomic databases such as UK Biobank and All of
Us. All cross-validated and/or weighted TMLE and OSE for the AIE k-point
interaction, as well as ATEs, conditional ATEs and functions thereof, are
implemented in the general purpose Julia package TMLE.jl. For high-throughput
applications in population genomics, we provide the open-source Nextflow
pipeline and software TarGene which integrates seamlessly with modern
high-performance and cloud computing platforms.


## AI 摘要

TarGene是一个用于群体遗传学研究的统一统计工作流程，采用半参数高效且双重稳健的估计方法，用于分析包括k点交互作用在内的遗传效应。它通过交叉验证和加权版本的TMLE和OSE估计平均交互效应（AIEs），并利用基于遗传相关性的方差估计校正数据依赖性。该方法适用于大规模基因组数据库（如UK Biobank和All of Us），能有效控制模型误设偏差和假发现率。相关功能已集成到Julia包TMLE.jl中，并提供了支持高性能和云计算的开源Nextflow流程TarGene，适用于高通量群体基因组学研究。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T22:01:17Z
- **目录日期**: 2025-05-21
