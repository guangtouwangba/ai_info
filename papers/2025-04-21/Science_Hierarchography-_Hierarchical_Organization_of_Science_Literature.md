# Science Hierarchography: Hierarchical Organization of Science Literature

**URL**: http://arxiv.org/abs/2504.13834v1

## 原始摘要

Scientific knowledge is growing rapidly, making it challenging to track
progress and high-level conceptual links across broad disciplines. While
existing tools like citation networks and search engines make it easy to access
a few related papers, they fundamentally lack the flexible abstraction needed
to represent the density of activity in various scientific subfields. We
motivate SCIENCE HIERARCHOGRAPHY, the goal of organizing scientific literature
into a high-quality hierarchical structure that allows for the categorization
of scientific work across varying levels of abstraction, from very broad fields
to very specific studies. Such a representation can provide insights into which
fields are well-explored and which are under-explored. To achieve the goals of
SCIENCE HIERARCHOGRAPHY, we develop a range of algorithms. Our primary approach
combines fast embedding-based clustering with LLM-based prompting to balance
the computational efficiency of embedding methods with the semantic precision
offered by LLM prompting. We demonstrate that this approach offers the best
trade-off between quality and speed compared to methods that heavily rely on
LLM prompting, such as iterative tree construction with LLMs. To better reflect
the interdisciplinary and multifaceted nature of research papers, our hierarchy
captures multiple dimensions of categorization beyond simple topic labels. We
evaluate the utility of our framework by assessing how effectively an LLM-based
agent can locate target papers using the hierarchy. Results show that this
structured approach enhances interpretability, supports trend discovery, and
offers an alternative pathway for exploring scientific literature beyond
traditional search methods. Code, data and demo:
$\href{https://github.com/JHU-CLSP/science-hierarchography}{https://github.com/JHU-CLSP/science-hierarchography}$


## AI 摘要

这篇论文提出了"科学层级图"(SCIENCE HIERARCHOGRAPHY)的概念，旨在将科学文献组织成多层次的分层结构，从广泛领域到具体研究进行系统分类。研究者开发了一种结合快速嵌入聚类和LLM提示的算法，在计算效率和语义准确性之间取得平衡。该方法能反映论文的多维分类特征，超越了简单主题标签。评估显示，基于LLM的智能体利用该层级结构能更有效地定位目标论文，提高了可解释性并支持趋势发现，为科学文献探索提供了传统搜索方法之外的新途径。相关代码和数据已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-21T23:01:07Z
- **目录日期**: 2025-04-21
