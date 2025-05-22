# #一图详解如何用大模型训练小模型##一图详解三种蒸馏技术#三人行，必有我师焉！模型也能通过“学习”不断进步。蒸馏就是一种能让小型AI模型（学生模型）向大型AI...

**URL**: https://weibo.com/6105753431/Pt0wybvEr

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%B8%80%E5%9B%BE%E8%AF%A6%E8%A7%A3%E5%A6%82%E4%BD%95%E7%94%A8%E5%A4%A7%E6%A8%A1%E5%9E%8B%E8%AE%AD%E7%BB%83%E5%B0%8F%E6%A8%A1%E5%9E%8B%23&amp;extparam=%23%E4%B8%80%E5%9B%BE%E8%AF%A6%E8%A7%A3%E5%A6%82%E4%BD%95%E7%94%A8%E5%A4%A7%E6%A8%A1%E5%9E%8B%E8%AE%AD%E7%BB%83%E5%B0%8F%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#一图详解如何用大模型训练小模型#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%B8%80%E5%9B%BE%E8%AF%A6%E8%A7%A3%E4%B8%89%E7%A7%8D%E8%92%B8%E9%A6%8F%E6%8A%80%E6%9C%AF%23&amp;extparam=%23%E4%B8%80%E5%9B%BE%E8%AF%A6%E8%A7%A3%E4%B8%89%E7%A7%8D%E8%92%B8%E9%A6%8F%E6%8A%80%E6%9C%AF%23" data-hide=""><span class="surl-text">#一图详解三种蒸馏技术#</span></a><br><br>三人行，必有我师焉！模型也能通过“学习”不断进步。<br><br>蒸馏就是一种能让小型AI模型（学生模型）向大型AI模型（教师模型）高效学习的技术。<br><br>接下来，让我们用一张图了解一下三种不同的蒸馏模式：<br><br>1. 软标签蒸馏（Soft-label distillation）<br><br>软标签蒸馏使用一个已训练且参数固定的教师模型和一个未训练的学生模型。<br><br>在软标签蒸馏中，学生模型通过匹配教师模型的预测概率分布（软标签）来学习。这要求学生模型在训练时能够访问教师模型对于每个训练样本所产生的概率输出。<br><br>2. ￼ 硬标签蒸馏（Hard-label distillation）<br><br>硬标签蒸馏中，学生模型直接学习教师模型的最终预测类别（即one-hot形式的硬标签）。<br><br>训练时，学生模型通过交叉熵损失优化自身输出，使其预测结果与教师模型的硬标签一致。<br><br>3. 协同蒸馏（Co-distillation）<br><br>协同蒸馏通常从多个未训练的模型开始，多个模型同时作为教师和学生。<br><br>它们在训练过程中彼此交换并学习对方的软标签概率分布，有时会共同接收真实硬标签的监督，从而共同进步。<img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1oc3nb0jkg30so0um7u4.gif" referrerpolicy="no-referrer"><br><br>

## AI 摘要

知识蒸馏是一种让小型AI模型(学生模型)向大型AI模型(教师模型)学习的技术。主要分为三种方法：(1)软标签蒸馏：学生模型学习教师模型的概率分布输出；(2)硬标签蒸馏：学生模型直接学习教师模型的最终预测类别；(3)协同蒸馏：多个模型互相学习对方的概率分布，共同提升。这些技术通过"师生"互动，实现了模型能力的有效迁移与压缩。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-22T12:03:35Z
- **目录日期**: 2025-05-22
