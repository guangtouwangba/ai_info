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

该研究提出了"科学层级图谱"(SCIENCE HIERARCHOGRAPHY)，旨在将科学文献组织成多层次分类体系，以追踪跨学科研究进展。研究团队开发了一种结合嵌入聚类与LLM提示的算法，在计算效率与语义精度间取得平衡。相比纯LLM方法，该方案在速度和质量上表现更优。层级结构能多维度分类论文，帮助识别研究热点与空白。测试表明，基于该结构的LLM代理能更有效地定位目标文献，提升可解释性并支持趋势发现，为传统检索方法提供了补充方案。相关代码和数据已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-22T02:28:25Z
- **目录日期**: 2025-04-22
