# PathOrchestra: A Comprehensive Foundation Model for Computational Pathology with Over 100 Diverse Clinical-Grade Tasks

**URL**: http://arxiv.org/abs/2503.24345v1

## 原始摘要

The complexity and variability inherent in high-resolution pathological
images present significant challenges in computational pathology. While
pathology foundation models leveraging AI have catalyzed transformative
advancements, their development demands large-scale datasets, considerable
storage capacity, and substantial computational resources. Furthermore,
ensuring their clinical applicability and generalizability requires rigorous
validation across a broad spectrum of clinical tasks. Here, we present
PathOrchestra, a versatile pathology foundation model trained via
self-supervised learning on a dataset comprising 300K pathological slides from
20 tissue and organ types across multiple centers. The model was rigorously
evaluated on 112 clinical tasks using a combination of 61 private and 51 public
datasets. These tasks encompass digital slide preprocessing, pan-cancer
classification, lesion identification, multi-cancer subtype classification,
biomarker assessment, gene expression prediction, and the generation of
structured reports. PathOrchestra demonstrated exceptional performance across
27,755 WSIs and 9,415,729 ROIs, achieving over 0.950 accuracy in 47 tasks,
including pan-cancer classification across various organs, lymphoma subtype
diagnosis, and bladder cancer screening. Notably, it is the first model to
generate structured reports for high-incidence colorectal cancer and
diagnostically complex lymphoma-areas that are infrequently addressed by
foundational models but hold immense clinical potential. Overall, PathOrchestra
exemplifies the feasibility and efficacy of a large-scale, self-supervised
pathology foundation model, validated across a broad range of clinical-grade
tasks. Its high accuracy and reduced reliance on extensive data annotation
underline its potential for clinical integration, offering a pathway toward
more efficient and high-quality medical services.


## AI 摘要

PathOrchestra是一个基于自监督学习的大规模病理基础模型，在30万张多中心病理切片上训练，涵盖20种组织和器官类型。该模型在112项临床任务（61个私有和51个公共数据集）上进行了严格验证，包括数字切片预处理、泛癌分类、病变识别、多癌亚型分类、生物标志物评估、基因表达预测和结构化报告生成。在27,755张全切片图像和9,415,729个ROI上表现出色，47项任务准确率超过0.950，尤其在结直肠癌和淋巴瘤结构化报告生成方面具有开创性。其高准确率和低标注依赖性展现了临床应用的巨大潜力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-01T13:08:57Z
- **目录日期**: 2025-04-01
