# Boosting Rust Unit Test Coverage through Hybrid Program Analysis and Large Language Models

**URL**: http://arxiv.org/abs/2506.09002v1

## 原始摘要

Unit testing is essential for ensuring software reliability and correctness.
Classic Search-Based Software Testing (SBST) methods and concolic
execution-based approaches for generating unit tests often fail to achieve high
coverage due to difficulties in handling complex program units, such as
branching conditions and external dependencies. Recent work has increasingly
utilized large language models (LLMs) to generate test cases, improving the
quality of test generation by providing better context and correcting errors in
the model's output. However, these methods rely on fixed prompts, resulting in
relatively low compilation success rates and coverage. This paper presents
PALM, an approach that leverages large language models (LLMs) to enhance the
generation of high-coverage unit tests. PALM performs program analysis to
identify branching conditions within functions, which are then combined into
path constraints. These constraints and relevant contextual information are
used to construct prompts that guide the LLMs in generating unit tests. We
implement the approach and evaluate it in 10 open-source Rust crates.
Experimental results show that within just two or three hours, PALM can
significantly improves test coverage compared to classic methods, with
increases in overall project coverage exceeding 50% in some instances and its
generated tests achieving an average coverage of 75.77%, comparable to human
effort (71.30%), highlighting the potential of LLMs in automated test
generation. We submitted 91 PALM-generated unit tests targeting new code. Of
these submissions, 80 were accepted, 5 were rejected, and 6 remain pending
review. The results demonstrate the effectiveness of integrating program
analysis with AI and open new avenues for future research in automated software
testing.


## AI 摘要

PALM是一种利用大语言模型(LLMs)提升单元测试生成的新方法。它通过程序分析识别函数分支条件，构建路径约束和上下文信息来生成更有效的测试提示。相比传统基于搜索或符号执行的方法，PALM能显著提高测试覆盖率，在Rust开源项目中平均覆盖率达75.77%，接近人工水平(71.30%)。实验显示，PALM仅需2-3小时就能使项目覆盖率提升超50%，且生成的91个测试中有80个被项目接受。该方法结合程序分析与AI，为自动化软件测试开辟了新途径。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-11T09:02:12Z
- **目录日期**: 2025-06-11
