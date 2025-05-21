# NExT-Search: Rebuilding User Feedback Ecosystem for Generative AI Search

**URL**: http://arxiv.org/abs/2505.14680v1

## 原始摘要

Generative AI search is reshaping information retrieval by offering
end-to-end answers to complex queries, reducing users' reliance on manually
browsing and summarizing multiple web pages. However, while this paradigm
enhances convenience, it disrupts the feedback-driven improvement loop that has
historically powered the evolution of traditional Web search. Web search can
continuously improve their ranking models by collecting large-scale,
fine-grained user feedback (e.g., clicks, dwell time) at the document level. In
contrast, generative AI search operates through a much longer search pipeline,
spanning query decomposition, document retrieval, and answer generation, yet
typically receives only coarse-grained feedback on the final answer. This
introduces a feedback loop disconnect, where user feedback for the final output
cannot be effectively mapped back to specific system components, making it
difficult to improve each intermediate stage and sustain the feedback loop. In
this paper, we envision NExT-Search, a next-generation paradigm designed to
reintroduce fine-grained, process-level feedback into generative AI search.
NExT-Search integrates two complementary modes: User Debug Mode, which allows
engaged users to intervene at key stages; and Shadow User Mode, where a
personalized user agent simulates user preferences and provides AI-assisted
feedback for less interactive users. Furthermore, we envision how these
feedback signals can be leveraged through online adaptation, which refines
current search outputs in real-time, and offline update, which aggregates
interaction logs to periodically fine-tune query decomposition, retrieval, and
generation models. By restoring human control over key stages of the generative
AI search pipeline, we believe NExT-Search offers a promising direction for
building feedback-rich AI search systems that can evolve continuously alongside
human feedback.


## AI 摘要

生成式AI搜索通过提供端到端答案改变了传统信息检索模式，但同时也削弱了基于用户反馈的持续优化机制。传统搜索依赖文档级细粒度反馈（如点击、停留时间）改进排序模型，而AI搜索的长流程（查询分解、文档检索、答案生成）仅能获得最终答案的粗粒度反馈，导致各环节优化困难。为此，研究提出NExT-Search新范式：通过"用户调试模式"让用户干预关键环节，利用"影子用户模式"模拟用户偏好；结合实时在线调整和离线模型更新，重建细粒度反馈循环，使AI搜索系统能随人类反馈持续进化。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T14:01:02Z
- **目录日期**: 2025-05-21
